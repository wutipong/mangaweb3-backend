package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func UpdateTags() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		log.Info().Str("item", m.Name).Msg("Populate tags.")
		_, err := meta.PopulateTags(context.Background(), m)
		if err != nil {
			log.Err(err).Msg("fails to populate tags.")
		}
	}
	return nil
}

func ScheduleUpdateTags() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Info().Msg("Update tags.")
		UpdateTags()
	})
}
