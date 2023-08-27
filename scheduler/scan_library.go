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
				break
			}
		}
		if found {
			continue
		}

		log.Info().
			Str("file", file).
			Msg("Creating metadata.")

		item, err := meta.NewItem(file)
		if err != nil {
			log.
				Error().
				AnErr("error", err).
				Msg("Failed to create meta data.")
		}

		err = meta.Write(context.Background(), item)
		if err != nil {
			log.Error().
				AnErr("error", err).
				Msg("Failed to write meta data.")
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

		log.Info().Str("file", m.Name).Msg("Deleting metadata.")
		if err := meta.Delete(context.Background(), m); err != nil {
			log.Error().
				Str("name", m.Name).
				AnErr("error", err).
				Msg("Failed to delete meta")
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
