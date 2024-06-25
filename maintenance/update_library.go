package maintenance

import (
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
)

func UpdateLibrary(client *ent.Client) {
	log.Info().Msg("Update metadata set.")
	ScanLibrary(client)
	log.Info().Msg("Update tag list.")
	UpdateTags(client)
	log.Info().Msg("Update missing thumbnails.")
	UpdateMissingThumbnail(client)
}
