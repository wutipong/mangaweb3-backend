package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/tag"
)

func RebuildTagThumbnail() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	allTags, err := tag.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, t := range allTags {
		log.Info().Str("tag", t.Name).Msg("Updating tag thumbnail.")
		for _, m := range allMeta {
			if true { // !slices.Contains(m.Tags, t.Name) {
				continue
			}

			t.Thumbnail = m.Thumbnail
			if err := tag.Write(context.Background(), t); err != nil {
				return err
			}

			break
		}
	}

	return nil
}

func ScheduleRebuildTagThumbnail() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Info().Msg("Rebuild tag thumbnails.")
		RebuildTagThumbnail()
	})
}
