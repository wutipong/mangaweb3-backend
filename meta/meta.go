package meta

import (
	"archive/zip"
	"bytes"
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
	"github.com/wutipong/mangaweb3-backend/tag"

	_ "golang.org/x/image/webp"
)

var client *ent.Client

func Init(c *ent.Client) {
	client = c
}

func NewItem(name string) (i *ent.Meta, err error) {
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

	PopulateTags(i)
	return
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

	const thumbnailSize = 200
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

func PopulateTags(m *ent.Meta) {
	m.Tags = tag.ParseTag(m.Name)
}
