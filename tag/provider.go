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

func Delete(ctx context.Context, t ent.Tag) error {
	return client.Tag.DeleteOne(&t).Exec(ctx)
}

func IsTagExist(ctx context.Context, name string) bool {
	count, err := client.Tag.Query().Where(tag.Name(name)).Count(ctx)
	if err != nil {
		return false
	}

	return count > 0
}

func Read(ctx context.Context, name string) (t *ent.Tag, err error) {
	return client.Tag.Query().Where(tag.Name(name)).Only(ctx)
}

func ReadAll(ctx context.Context) (tags []*ent.Tag, err error) {
	return client.Tag.Query().All(ctx)

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
