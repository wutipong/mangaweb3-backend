package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/tag"
)

func UpdateTags() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	tagSet := make(map[string]bool)
	for _, m := range allMeta {
		tags := tag.ParseTag(m.Name)
		for _, t := range tags {
			tagSet[t] = true
		}
	}

	allTag, err := tag.ReadAll(context.Background())
	if err != nil {
		log.Error().
			AnErr("error", err).
			Msg("Cannot read metadata.")
		return err
	}

	allTagSet := make(map[string]bool)
	for _, t := range allTag {
		allTagSet[t.Name] = true
	}

	findMetaWithTag := func(tag string) meta.Meta {
		for _, m := range allMeta {
			for _, t := range m.Tags {
				if t == tag {
					return m
				}
			}
		}

		return meta.Meta{}
	}

	for tagStr := range tagSet {
		if _, found := allTagSet[tagStr]; !found {
			t := tag.NewTag(tagStr)
			m := findMetaWithTag(tagStr)
			t.Thumbnail = m.Thumbnail

			err = tag.Write(context.Background(), t)

			if err != nil {
				return err
			}
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
