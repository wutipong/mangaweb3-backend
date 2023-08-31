package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/tag"
)

func RebuildTagThumbnail() error {
	allTags, err := tag.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, t := range allTags {
		log.Info().Str("tag", t.Name).Msg("Updating tag thumbnail.")
		m, err := t.QueryMeta().
			First(context.TODO())
		if err != nil {
			log.Err(err).Msg("update fails")
			continue
		}

		t.Thumbnail = m.Thumbnail
		t.Update().SetThumbnail(t.Thumbnail).Save(context.TODO())
	}

	return nil
}

func ScheduleRebuildTagThumbnail() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Info().Msg("Rebuild tag thumbnails.")
		RebuildTagThumbnail()
	})
}
