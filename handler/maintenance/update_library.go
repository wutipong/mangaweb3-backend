package maintenance

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/maintenance"
)

type UpdateLibraryResponse struct {
	Result bool `json:"result"`
}

const (
	PathUpdateLibrary = "/maintenance/update_library"
)

// @Success      200  {object}  maintenance.UpdateLibraryResponse
// @Failure      500  {object}  errors.Error
// @Router /maintenance/update_library [get]
func UpdateLibraryHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Info().Msg("Update library")

	go maintenance.UpdateLibrary(r.Context())

	response := UpdateLibraryResponse{
		Result: true,
	}

	handler.WriteResponse(w, response)
}
