package meta

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"sort"
	"time"

	"github.com/disintegration/imaging"
	"github.com/facette/natsort"
	"github.com/wutipong/mangaweb3-backend/data"
	"github.com/wutipong/mangaweb3-backend/ent"
	tag_util "github.com/wutipong/mangaweb3-backend/tag"

	"golang.org/x/exp/slices"
	_ "golang.org/x/image/webp"
)

func NewItem(ctx context.Context, client *ent.Client, name string) (i *ent.Meta, err error) {
	createTime := time.Now()
	if t, e := data.GetLastModifiedTime(ctx, name); e == nil {
		createTime = t
	}

	i = &ent.Meta{
		Name:       name,
		CreateTime: createTime,
		Favorite:   false,
	}

	if err = GenerateImageIndices(ctx, i); err != nil {
		return
	}

	GenerateThumbnail(ctx, i, 0, CropDetails{})

	return client.Meta.Create().
		SetName(i.Name).
		SetCreateTime(i.CreateTime).
		SetFavorite(i.Favorite).
		SetFileIndices(i.FileIndices).
		SetRead(false).
		SetThumbnail(i.Thumbnail).Save(ctx)
}

type CropDetails struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func GenerateThumbnail(ctx context.Context, m *ent.Meta, fileIndex int, details CropDetails) error {
	if len(m.FileIndices) == 0 {
		return fmt.Errorf("file list is empty")
	}

	children, err := data.ListObject(ctx, m.Name)
	if err != nil {
		return err
	}

	stream, err := data.GetObject(ctx, children[m.FileIndices[fileIndex]])
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

func GenerateImageIndices(ctx context.Context, m *ent.Meta) error {
	children, err := data.ListObject(ctx, m.Name)
	if err != nil {
		return err
	}

	type fileIndexPair struct {
		Index    int
		FileName string
	}

	var fileNames []fileIndexPair
	for i, f := range children {
		if filter(f) {
			fileNames = append(fileNames,
				fileIndexPair{
					i, f,
				})
		}
	}

	sort.Slice(fileNames, func(i, j int) bool {
		return natsort.Compare(fileNames[i].FileName, fileNames[j].FileName)
	})

	m.FileIndices = make([]int, len(fileNames))
	for i, p := range fileNames {
		m.FileIndices[i] = p.Index
	}

	return nil
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
