package container

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/facette/natsort"
	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/ent"
)

type ZipContainer struct {
	Meta *ent.Meta
}

func (c *ZipContainer) ListItems(ctx context.Context) (names []string, err error) {
	m := c.Meta

	fullpath := configuration.Get().DataPath + string(os.PathSeparator) + m.Name

	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return
	}
	defer r.Close()

	names = make([]string, len(m.FileIndices))
	for i, f := range m.FileIndices {
		names[i] = r.File[f].Name
	}

	return
}

func (c *ZipContainer) OpenItem(ctx context.Context, index int) (reader io.ReadCloser, name string, err error) {
	dataPath := configuration.Get().DataPath

	fullpath := filepath.Join(dataPath, c.Meta.Name)

	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return
	}

	defer r.Close()

	if index >= len(c.Meta.FileIndices) {
		err = fmt.Errorf("invalid item")
		return
	}

	zf := r.File[c.Meta.FileIndices[index]]

	if zf == nil {
		err = fmt.Errorf("file not found : %v", index)
		return
	}

	name = zf.Name
	reader, err = zf.Open()
	if err != nil {
		return
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		return
	}

	reader = io.NopCloser(bytes.NewBuffer(content))

	return
}

func (c *ZipContainer) PopulateImageIndices(ctx context.Context) error {
	m := c.Meta

	fullpath := configuration.Get().DataPath + string(os.PathSeparator) + m.Name

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
		if isValidImageFile(f.Name) {
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

func (c *ZipContainer) Download(ctx context.Context) (reader io.ReadCloser, filename string, err error) {
	fullpath := filepath.Join(configuration.Get().DataPath, c.Meta.Name)
	reader, err = os.Open(fullpath)
	filename = filepath.Base(c.Meta.Name)

	return
}
