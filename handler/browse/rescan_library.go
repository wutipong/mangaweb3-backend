package browse

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/scheduler"
)

type rescanLibraryResponse struct {
	Result bool `json:"result"`
}

const (
	PathRescanLibrary = "/browse/rescan_library"
)

// @Success      200  {object}  browse.rescanLibraryResponse
// @Failure      500  {object}  errors.Error
// @Router /browse/rescan_library [get]
func RescanLibraryHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Rescan library")

	scheduler.ScheduleScanLibrary(handler.EntClient())

	response := rescanLibraryResponse{
		Result: true,
	}

	handler.WriteResponse(w, response)
}
