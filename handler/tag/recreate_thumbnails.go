package tag

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/database"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/maintenance"
)

type recreateThumbnailsResponse struct {
	Result bool `json:"result"`
}

const (
	PathRecreateThumbnails = "/tag/recreate_thumbnails"
)

// @Success      200  {object}  browse.recreateThumbnailsResponse
// @Failure      500  {object}  errors.Error
// @Router /tag/recreate_thumbnails [get]
func RecreateThumbnailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Info().Msg("Rescan library")

	client := database.CreateEntClient()
	defer client.Close()
	go maintenance.RebuildTagThumbnail(client)

	response := recreateThumbnailsResponse{
		Result: true,
	}

	handler.WriteResponse(w, response)
}
