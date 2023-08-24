package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/scheduler"
)

type rescanLibraryResponse struct {
	Result bool `json:"result"`
}

// @Success      200  {object}  handler.rescanLibraryResponse
// @Failure      500  {object}  errors.Error
// @Router /rescan_library [get]
func RescanLibraryHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Info().Msg("Rescan library")

	scheduler.ScheduleScanLibrary()

	response := rescanLibraryResponse{
		Result: true,
	}

	WriteResponse(w, response)
}
