package browse

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/scheduler"
)

type recreateThumbnailsResponse struct {
	Result bool `json:"result"`
}

const (
	PathRecreateThumbnails = "/browse/recreate_thumbnails"
)

// @Success      200  {object}  browse.recreateThumbnailsResponse
// @Failure      500  {object}  errors.Error
// @Router /browse/recreate_thumbnails [get]
func RecreateThumbnailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Info().Msg("Rescan library")

	scheduler.ScheduleRebuildThumbnail()

	response := recreateThumbnailsResponse{
		Result: true,
	}

	handler.WriteResponse(w, response)
}
