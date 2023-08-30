package meta

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
)

func IsItemExist(ctx context.Context, name string) bool {
	count, err := client.Meta.Query().Where(meta.Name(name)).Count(ctx)
	if err != nil {
		return false
	}

	return count > 0
}
func Write(ctx context.Context, i *ent.Meta) error {
	return client.Meta.Create().
		SetName(i.Name).
		SetCreateTime(i.CreateTime).
		SetFavorite(i.Favorite).
		SetFileIndices(i.FileIndices).
		SetThumbnail(i.Thumbnail).
		SetRead(i.Read).
		OnConflict(sql.ConflictColumns(meta.FieldName)).
		UpdateNewValues().Exec(ctx)
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
