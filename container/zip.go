package container

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/ent"
)

type ZipContainer struct {
	Meta *ent.Meta
}

func (c *ZipContainer) GetSize(ctx context.Context) int {
	return len(c.Meta.FileIndices)
}

func (c *ZipContainer) OpenItem(ctx context.Context, index int) (reader io.ReadCloser, name string, err error) {
	dataPath := configuration.Get().DataPath

	fullpath := filepath.Join(dataPath, c.Meta.Name)

	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return
	}

	defer r.Close()

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
