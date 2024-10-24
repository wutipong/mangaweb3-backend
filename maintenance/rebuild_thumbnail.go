package maintenance

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func RebuildThumbnail(client *ent.Client) error {
	allMeta, err := meta.ReadAll(context.Background(), client)
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		e := meta.GenerateThumbnail(context.Background(), m, 0, meta.CropDetails{})
		log.Info().Str("name", m.Name).Msg("Create new thumbnail")
		if e != nil {
			log.Error().Str("item", m.Name).AnErr("error", e).Msg("Fails to create thumbnail.")
			continue
		}

		meta.Write(context.Background(), client, m)
	}

	return nil
}
