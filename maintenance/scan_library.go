package maintenance

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func ScanLibrary(client *ent.Client) error {
	allMeta, err := meta.ReadAll(context.Background(), client)
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
			if m.Name == file.Name {
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
			Str("file", file.Name).
			Str("type", file.Type.String()).
			Msg("Creating metadata.")

		if item, err := meta.Read(context.Background(), client, file.Name); err == nil {
			item.Active = true
			if err := meta.Write(context.Background(), client, item); err != nil {
				log.Error().
					Str("name", item.Name).
					AnErr("error", err).
					Msg("Failed to re-activate meta")
			}
		} else {
			item, err := meta.NewItem(context.Background(), client, file.Name, file.Type)
			if err != nil {
				log.
					Error().
					AnErr("error", err).
					Msg("Failed to create meta data.")

				continue
			}

			_, err = meta.PopulateTags(context.Background(), client, item)
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
			if m.Name == file.Name {
				found = true
				break
			}
		}
		if found {
			continue
		}

		log.Info().Str("file", m.Name).Msg("Inactivate metadata.")
		m.Active = false

		if err := meta.Write(context.Background(), client, m); err != nil {
			log.Error().
				Str("name", m.Name).
				AnErr("error", err).
				Msg("Failed to inactivate meta")
		}
	}

	return nil
}
