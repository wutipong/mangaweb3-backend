package view

import (
	"archive/zip"
	"path/filepath"

	"github.com/wutipong/mangaweb3-backend/config"
	"github.com/wutipong/mangaweb3-backend/ent"
)

type Page struct {
	Index int
	Name  string
}

func ListPages(m *ent.Meta) (pages []Page, err error) {
	if len(m.FileIndices) == 0 {
		return
	}

	c := config.Get()

	fullpath := filepath.Join(c.DataPath, m.Name)
	r, err := zip.OpenReader(fullpath)
	if err != nil {
		return
	}

	defer r.Close()

	pages = make([]Page, len(m.FileIndices))
	for i, f := range m.FileIndices {
		pages[i] = Page{
			Name:  r.File[f].Name,
			Index: i,
		}
	}

	return
}
