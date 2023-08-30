package meta

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/facette/natsort"
	"github.com/wutipong/mangaweb3-backend/ent"
	tag_util "github.com/wutipong/mangaweb3-backend/tag"

	"golang.org/x/exp/slices"
	_ "golang.org/x/image/webp"
)

var client *ent.Client

func Init(c *ent.Client) {
	client = c
}

func NewItem(ctx context.Context, name string) (i *ent.Meta, err error) {
	createTime := time.Now()

	if stat, e := fs.Stat(os.DirFS(BaseDirectory), name); e == nil {
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

	if err = GenerateThumbnail(i, 0); err != nil {
		return
	}

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

	fullpath := filepath.Join(BaseDirectory, m.Name)

	reader, err = os.Open(fullpath)
	return
}

func GenerateThumbnail(m *ent.Meta, fileIndex int) error {
	mutex := new(sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	fullpath := filepath.Join(BaseDirectory, m.Name)

	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return err
	}
	defer r.Close()

	if len(m.FileIndices) == 0 {
		return fmt.Errorf("file list is empty")
	}

	stream, err := r.File[m.FileIndices[fileIndex]].Open()
	if err != nil {
		return err
	}

	defer stream.Close()

	img, err := imaging.Decode(stream, imaging.AutoOrientation(true))
	if err != nil {
		return err
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

	fullpath := BaseDirectory + string(os.PathSeparator) + m.Name

	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return err
	}
	defer r.Close()

	type fileIndexPair struct {
		Index    int
		FileName string
	}

	var fileNames []fileIndexPair
	for i, f := range r.File {
		if filter(f.Name) {
			fileNames = append(fileNames,
				fileIndexPair{
					i, f.Name,
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

func PopulateTags(ctx context.Context, m *ent.Meta) (out *ent.Meta, err error) {
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
		if temp, err := tag_util.Read(ctx, t); err != nil {
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
