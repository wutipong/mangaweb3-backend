package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func UpdateMissingThumbnail() error {

	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		if len(m.Thumbnail) != 0 {
			continue
		}
		e := meta.GenerateThumbnail(m, 0)

		log.Info().
			Str("name", m.Name).
			Msg("generating new thumbnail.")
		if e != nil {
			log.Error().
				Str("name", m.Name).
				AnErr("error", err).
				Msg("Failed to generate thumbnail.")
			continue
		}

		meta.Write(context.Background(), m)
	}

	return nil
}

func ScheduleUpdateMissingThumbnail() {
	scheduler.Every(1).Hour().Do(func() {
		log.Info().Msg("Updating missing thumbnail.")
		UpdateMissingThumbnail()
	})
}
