package container

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
)

type Container interface {
	ListItems(ctx context.Context) (names []string, err error)
	OpenItem(ctx context.Context, index int) (reader io.ReadCloser, name string, err error)
	PopulateImageIndices(ctx context.Context) error
}

func GuessContainerType(ctx context.Context, name string, info fs.FileInfo) (t meta.ContainerType, valid bool) {
	valid = false

	if !isValidContainerName(info.Name()) {
		return
	}

	if info.IsDir() {
		///TODO: add support for directory type

		t = meta.ContainerTypeDirectory
		valid = false

		return
	}

	ext := strings.ToLower(filepath.Ext(info.Name()))

	if ext == ".zip" || ext == ".cbz" {
		t = meta.ContainerTypeZip
		valid = true

		return
	}

	return
}

func CreateContainer(m *ent.Meta) (c Container, err error) {
	switch m.ContainerType {
	case meta.ContainerTypeZip:
		c = &ZipContainer{
			Meta: m,
		}

		return

	default:
		err = fmt.Errorf("invalid container type")
		return
	}
}

func isValidContainerName(name string) bool {
	if strings.HasPrefix(name, ".") {
		return false
	}

	return true
}

func isValidImageFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	if ext == ".jpeg" {
		return true
	}
	if ext == ".jpg" {
		return true
	}
	if ext == ".png" {
		return true
	}
	if ext == ".webp" {
		return true
	}
	return false
}
