package maintenance

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func UpdateTags(client *ent.Client) error {
	allMeta, err := meta.ReadAll(context.Background(), client)
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		log.Info().Str("item", m.Name).Msg("Populate tags.")
		_, _, err := meta.PopulateTags(context.Background(), client, m)
		if err != nil {
			log.Err(err).Msg("fails to populate tags.")
		}
	}
	return nil
}
