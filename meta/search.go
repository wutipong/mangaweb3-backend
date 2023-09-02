package meta

import (
	"context"

	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
	"github.com/wutipong/mangaweb3-backend/ent/tag"
)

type SearchField string
type SortField string
type SortOrder string

const (
	SearchFieldName     = SearchField("name")
	SearchFieldFavorite = SearchField("favorite")
	SearchFieldTag      = SearchField("tag")

	SortFieldName       = SortField("name")
	SortFieldCreateTime = SortField("createTime")

	SortOrderAscending  = SortOrder("ascending")
	SortOrderDescending = SortOrder("descending")
)

type SearchCriteria struct {
	Field SearchField
	Value interface{}
}

func SearchFilter(ctx context.Context,
	name string,
	favoriteOnly bool,
	searchTag string,
	sortBy SortField,
	sortOrder SortOrder,
	page int,
	itemPerPage int) (query *ent.MetaQuery, err error) {

	if searchTag != "" {
		t, e := client.Tag.Query().Where(tag.Name(searchTag)).Only(ctx)
		if e != nil {
			err = e
			return
		}

		query = t.QueryMeta()
	} else {
		query = client.Meta.Query()
	}

	predicates := []predicate.Meta{
		meta.Active(true),
	}

	if name != "" {
		predicates = append(predicates, meta.NameContains(name))
	}

	if favoriteOnly {
		predicates = append(predicates, meta.Favorite(true))
	}

	query = query.Where(predicates...)

	field := ""
	switch sortBy {
	case SortFieldName:
		field = meta.FieldName
	case SortFieldCreateTime:
		field = meta.FieldCreateTime
	}

	switch sortOrder {
	case SortOrderAscending:
		query = query.Order(ent.Asc(string(field)))
	case SortOrderDescending:
		query = query.Order(ent.Desc(string(field)))
	}

	if itemPerPage > 0 {
		query = query.Limit(itemPerPage).Offset(itemPerPage * page)
	}

	return
}

func SearchItems(ctx context.Context,
	name string,
	favoriteOnly bool,
	searchTag string,
	sortBy SortField,
	sortOrder SortOrder,
	page int,
	itemPerPage int,
) (items []*ent.Meta, err error) {

	query, err := SearchFilter(ctx, name, favoriteOnly, searchTag, sortBy, sortOrder, page, itemPerPage)
	if err != nil {
		return
	}

	return query.All(ctx)
}

func CountItems(ctx context.Context,
	name string,
	favoriteOnly bool,
	searchTag string,
	sortBy SortField,
	sortOrder SortOrder,
) (count int, err error) {

	query, err := SearchFilter(ctx, name, favoriteOnly, searchTag, sortBy, sortOrder, 0, 0)
	if err != nil {
		return
	}

	return query.Count(ctx)
}
