package maintenance

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/database"
)

func UpdateLibrary(ctx context.Context) {
	client := database.CreateEntClient()
	defer client.Close()

	log.Info().Msg("Update metadata set.")
	ScanLibrary(ctx, client)
}
