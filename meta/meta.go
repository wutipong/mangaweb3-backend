package meta

import (
	"bytes"
	"context"
	"image"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/container"
	"github.com/wutipong/mangaweb3-backend/ent"
	tag_util "github.com/wutipong/mangaweb3-backend/tag"

	"golang.org/x/exp/slices"
	_ "golang.org/x/image/webp"
)

func NewItem(ctx context.Context, client *ent.Client, name string) (i *ent.Meta, err error) {
	createTime := time.Now()

	c := configuration.Get()
	if stat, e := fs.Stat(os.DirFS(c.DataPath), name); e == nil {
		createTime = stat.ModTime()
	}

	i = &ent.Meta{
		Name:       name,
		CreateTime: createTime,
		Favorite:   false,
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

	const thumbnailSize = 360
	if img.Bounds().Dx() > thumbnailSize {
		resized := imaging.Resize(img, thumbnailSize, 0, imaging.MitchellNetravali)
		img = resized
	}

	buffer := bytes.Buffer{}
	imaging.Encode(&buffer, img, imaging.JPEG, imaging.JPEGQuality(75))

	m.Thumbnail = buffer.Bytes()

	return nil
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
