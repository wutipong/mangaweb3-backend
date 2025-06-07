package browse

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/database"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/maintenance"
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
func RescanLibraryHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Info().Msg("Rescan library")

	client := database.CreateEntClient()
	defer client.Close()

	go maintenance.ScanLibrary(r.Context(), client)

	response := rescanLibraryResponse{
		Result: true,
	}

	handler.WriteResponse(w, response)
}
