package view

import (
	"context"

	"github.com/wutipong/mangaweb3-backend/data"
	"github.com/wutipong/mangaweb3-backend/ent"
)

type Page struct {
	Index int
	Name  string
}

func ListPages(ctx context.Context, m *ent.Meta) (pages []Page, err error) {
	if len(m.FileIndices) == 0 {
		return
	}

	children, err := data.ListObject(ctx, m.Name)

	pages = make([]Page, len(m.FileIndices))
	for i, f := range m.FileIndices {
		pages[i] = Page{
			Name:  children[f],
			Index: i,
		}
	}

	return
}
