package meta

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
)

func IsItemExist(ctx context.Context, client *ent.Client, name string) bool {
	count, err := client.Meta.Query().Where(meta.Name(name)).Count(ctx)
	if err != nil {
		return false
	}

	return count > 0
}
func Write(ctx context.Context, client *ent.Client, m *ent.Meta) error {
	return client.Meta.Create().
		SetName(m.Name).
		SetCreateTime(m.CreateTime).
		SetFavorite(m.Favorite).
		SetFileIndices(m.FileIndices).
		SetThumbnail(m.Thumbnail).
		SetRead(m.Read).
		SetActive(m.Active).
		OnConflict(sql.ConflictColumns(meta.FieldName)).
		UpdateNewValues().Exec(ctx)
}

func Read(ctx context.Context, client *ent.Client, name string) (m *ent.Meta, err error) {
	return client.Meta.Query().Where(meta.Name(name)).Only(ctx)
}

func ReadAll(ctx context.Context, client *ent.Client) (items []*ent.Meta, err error) {
	return client.Meta.Query().Where(meta.Active(true)).All(ctx)
}
