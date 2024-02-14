package tag

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/tag"
)

var client *ent.Client

func Init(c *ent.Client) {
	client = c
}

func IsTagExist(ctx context.Context, name string) bool {
	count, err := client.Tag.Query().Where(tag.Name(name)).Count(ctx)
	if err != nil {
		return false
	}

	return count > 0
}

func Read(ctx context.Context, name string) (t *ent.Tag, err error) {
	return client.Tag.Query().Where(tag.Name(name)).First(ctx)
}

func ReadAll(ctx context.Context) (tags []*ent.Tag, err error) {
	return client.Tag.Query().Order(tag.ByName()).All(ctx)
}

func ReadPage(ctx context.Context, favoriteOnly bool, search string,
	page int, itemPerPage int) (tags []*ent.Tag, err error) {
	query := client.Tag.Query().
		Offset(page * itemPerPage).
		Limit(itemPerPage).
		Order(tag.ByName())

	if favoriteOnly {
		query = query.Where(tag.Favorite(favoriteOnly))
	}

	if search != "" {
		query = query.Where(tag.NameContainsFold(search))
	}

	return query.All(ctx)
}

func Count(ctx context.Context, favoriteOnly bool) (count int, err error) {
	query := client.Tag.Query()
	if favoriteOnly {
		query = query.Where(tag.Favorite(favoriteOnly))
	}

	return query.Count(ctx)
}

func Write(ctx context.Context, t *ent.Tag) error {
	return client.Tag.Create().
		SetName(t.Name).
		SetHidden(t.Hidden).
		SetFavorite(t.Favorite).
		SetThumbnail(t.Thumbnail).
		OnConflict(sql.ConflictColumns(tag.FieldName)).
		UpdateNewValues().
		Exec(ctx)
}
