package meta

import (
	"context"

	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
)

func IsItemExist(ctx context.Context, name string) bool {
	count, err := client.Meta.Query().Where(meta.Name(name)).Count(ctx)
	if err != nil {
		return false
	}

	return count > 0
}
func Write(ctx context.Context, i *ent.Meta) (m *ent.Meta, err error) {
	m, err = client.Meta.Create().
		SetName(i.Name).
		SetCreateTime(i.CreateTime).
		SetFavorite(i.Favorite).
		SetFileIndices(i.FileIndices).
		SetThumbnail(i.Thumbnail).
		SetRead(i.Read).
		SetTags(i.Tags).
		Save(ctx)

	return
}

func Delete(ctx context.Context, i *ent.Meta) error {
	return client.Meta.DeleteOne(i).Exec(ctx)
}

func Read(ctx context.Context, name string) (i *ent.Meta, err error) {
	return client.Meta.Query().Where(meta.Name(name)).Only(ctx)
}

func ReadAll(ctx context.Context) (items []*ent.Meta, err error) {
	return client.Meta.Query().All(ctx)
}

func Search(ctx context.Context,
	criteria []SearchCriteria,
	sort SortField,
	order SortOrder,
	pageSize int,
	page int) (items []*ent.Meta, err error) {

	predicates := populatePredicates(criteria, []predicate.Meta{})

	return client.Meta.Query().Where(predicates...).Limit(pageSize).Offset(pageSize * page).All(ctx)

}

func populatePredicates(criteria []SearchCriteria, predicates []predicate.Meta) []predicate.Meta {
	for _, c := range criteria {
		switch c.Field {
		case SearchFieldName:
			predicates = append(predicates, meta.NameContains(c.Value.(string)))
		case SearchFieldFavorite:
			predicates = append(predicates, meta.Favorite(c.Value.(bool)))

		case SearchFieldTag:
			///TODO: Field tags is broken.
			// predicates = append(predicates, meta. Favorite(c.Value.(bool)))
		}
	}
	return predicates
}
func Count(ctx context.Context, criteria []SearchCriteria) (count int, err error) {
	predicates := populatePredicates(criteria, []predicate.Meta{})

	return client.Meta.Query().Where(predicates...).Count(ctx)
}
