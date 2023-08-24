package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func RebuildThumbnail() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		e := m.GenerateThumbnail(0)
		log.Info().Str("Generating new thumbnail for", m.Name)
		if e != nil {
			log.Error().Str("item", m.Name).AnErr("error", e).Msg("Fails to create thumbnail.")
			continue
		}

		meta.Write(context.Background(), m)
	}

	return nil
}

func ScheduleRebuildThumbnail() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Print("Force updating thumbnail")
		RebuildThumbnail()
	})
}
