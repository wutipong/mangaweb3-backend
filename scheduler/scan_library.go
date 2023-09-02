package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func ScanLibrary() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	files, err := meta.ListDir("")
	if err != nil {
		return err
	}

	for _, file := range files {
		found := false
		for _, m := range allMeta {
			if m.Name == file {
				found = true

				if m.Active {
					break
				}
			}
		}
		if found {
			continue
		}

		log.Info().
			Str("file", file).
			Msg("Creating metadata.")

		if item, err := meta.Read(context.Background(), file); err == nil {
			item.Active = true
			if err := meta.Write(context.Background(), item); err != nil {
				log.Error().
					Str("name", item.Name).
					AnErr("error", err).
					Msg("Failed to re-activate meta")
			}
		} else {
			item, err := meta.NewItem(context.Background(), file)
			if err != nil {
				log.
					Error().
					AnErr("error", err).
					Msg("Failed to create meta data.")
			}

			_, err = meta.PopulateTags(context.Background(), item)
			if err != nil {
				log.Error().
					AnErr("error", err).
					Msg("Failed to write meta data.")
			}
		}
	}

	for _, m := range allMeta {
		found := false
		for _, file := range files {
			if m.Name == file {
				found = true
				break
			}
		}
		if found {
			continue
		}

		log.Info().Str("file", m.Name).Msg("Inactivate metadata.")
		m.Active = false

		if err := meta.Write(context.Background(), m); err != nil {
			log.Error().
				Str("name", m.Name).
				AnErr("error", err).
				Msg("Failed to inactivate meta")
		}
	}

	return nil
}

func ScheduleScanLibrary() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Info().Msg("Scanning Library.")
		ScanLibrary()
	})
}
