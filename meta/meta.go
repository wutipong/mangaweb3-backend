package meta

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/container"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	tag_util "github.com/wutipong/mangaweb3-backend/tag"

	"golang.org/x/exp/slices"
	_ "golang.org/x/image/webp"
)

const (
	CACHE_LOCATION             = "cache"
	META_THUMB_LOCATION        = "meta"
	THUMBNAIL_FILENAME_PATTERN = "%d.jpg"

	THUMBNAIL_HEIGHT = 510
)

func NewItem(ctx context.Context, client *ent.Client, name string, ct meta.ContainerType) (i *ent.Meta, err error) {
	createTime := time.Now()

	c := configuration.Get()
	if stat, e := fs.Stat(os.DirFS(c.DataPath), name); e == nil {
		createTime = stat.ModTime()
	}

	i = &ent.Meta{
		Name:          name,
		CreateTime:    createTime,
		Favorite:      false,
		ContainerType: ct,
	}

	if err = GenerateImageIndices(i); err != nil {
		return
	}

	GenerateThumbnail(i, 0, CropDetails{})

	return client.Meta.Create().
		SetName(i.Name).
		SetCreateTime(i.CreateTime).
		SetFavorite(i.Favorite).
		SetFileIndices(i.FileIndices).
		SetRead(false).
		SetContainerType(ct).
		SetThumbnail(i.Thumbnail).Save(ctx)
}

func Open(m *ent.Meta) (reader io.ReadCloser, err error) {
	mutex := new(sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	c := configuration.Get()

	fullpath := filepath.Join(c.DataPath, m.Name)

	reader, err = os.Open(fullpath)
	return
}

type CropDetails struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func GenerateThumbnail(m *ent.Meta, fileIndex int, details CropDetails) error {
	mutex := new(sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	c, err := container.CreateContainer(m)
	if err != nil {
		return err
	}
	stream, _, err := c.OpenItem(context.Background(), fileIndex)
	if err != nil {
		return err
	}

	defer stream.Close()

	img, err := imaging.Decode(stream, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	if details.Width > 0 && details.Height > 0 {
		img = imaging.Crop(img, image.Rectangle{
			Min: image.Point{
				X: details.X,
				Y: details.Y,
			},
			Max: image.Point{
				X: details.X + details.Width,
				Y: details.Y + details.Height,
			},
		})
	}

	const thumbnailHeight = 510
	if img.Bounds().Dy() > thumbnailHeight {
		resized := imaging.Resize(img, 0, thumbnailHeight, imaging.MitchellNetravali)
		img = resized
	}

	buffer := bytes.Buffer{}
	imaging.Encode(&buffer, img, imaging.JPEG, imaging.JPEGQuality(75))

	m.Thumbnail = buffer.Bytes()

	return nil
}

func CreateThumbnail(m *ent.Meta) (thumbnail image.Image, err error) {
	mutex := new(sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	c, err := container.CreateContainer(m)
	if err != nil {
		return
	}

	stream, _, err := c.OpenItem(context.Background(), m.ThumbnailIndex)
	if err != nil {
		return
	}

	defer stream.Close()

	img, err := imaging.Decode(stream, imaging.AutoOrientation(true))
	if err != nil {
		return
	}

	if m.ThumbnailWidth > 0 && m.ThumbnailHeight > 0 {
		img = imaging.Crop(img, image.Rectangle{
			Min: image.Point{
				X: m.ThumbnailX,
				Y: m.ThumbnailY,
			},
			Max: image.Point{
				X: m.ThumbnailX + m.ThumbnailWidth,
				Y: m.ThumbnailY + m.ThumbnailHeight,
			},
		})
	}

	if img.Bounds().Dy() > THUMBNAIL_HEIGHT {
		resized := imaging.Resize(img, 0, THUMBNAIL_HEIGHT, imaging.MitchellNetravali)
		img = resized
	}

	thumbnail = img
	return
}

func CreateThumbnailPath(id int) string {
	return filepath.Join(CACHE_LOCATION, META_THUMB_LOCATION, fmt.Sprintf(THUMBNAIL_FILENAME_PATTERN, id))
}

func GetThumbnailBytes(m *ent.Meta) (thumbnail []byte, err error) {
	thumbfile := CreateThumbnailPath(m.ID)
	file, err := os.Open(thumbfile)
	buffer := bytes.Buffer{}

	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(filepath.Dir(thumbfile), fs.ModePerm)
		img, e := CreateThumbnail(m)
		if e != nil {
			err = e
			return
		}

		log.Debug().AnErr("error", imaging.Save(img, thumbfile, imaging.JPEGQuality(75))).Msg("save image")
		imaging.Encode(&buffer, img, imaging.JPEG, imaging.JPEGQuality(75))
		err = nil
	} else {
		io.Copy(&buffer, file)
	}

	thumbnail = bytes.Clone(buffer.Bytes())

	return
}

func DeleteThumbnail(m *ent.Meta) error {
	thumbfile := CreateThumbnailPath(m.ID)
	err := os.Remove(thumbfile)

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}

func GenerateImageIndices(m *ent.Meta) error {
	mutex := new(sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	c, err := container.CreateContainer(m)
	if err != nil {
		return err
	}

	return c.PopulateImageIndices(context.Background())
}

func PopulateTags(ctx context.Context, client *ent.Client, m *ent.Meta) (out *ent.Meta, err error) {
	tagStrs := tag_util.ParseTag(m.Name)
	currentTags, _ := m.QueryTags().All(ctx)

	newTags := make([]*ent.Tag, 0)
	for _, t := range tagStrs {

		if slices.ContainsFunc(currentTags, func(tag *ent.Tag) bool {
			return tag.Name == t
		}) {
			continue
		}

		var tag *ent.Tag
		if temp, err := tag_util.Read(ctx, client, t); err != nil {
			tag = &ent.Tag{
				Name: t,
			}

			tag, _ = client.Tag.Create().
				SetName(tag.Name).
				SetFavorite(tag.Favorite).
				SetHidden(tag.Hidden).
				SetThumbnail(m.Thumbnail).
				Save(ctx)

		} else {
			tag = temp
		}
		newTags = append(newTags, tag)
	}

	m, _ = m.Update().
		AddTags(newTags...).
		Save(ctx)

	out = m

	return
}
