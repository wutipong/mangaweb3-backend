package browse

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
)

const (
	PathThumbnail = "/browse/thumbnail"
)

// @Param name query string true "name of the item"
// @Success      200  {body}  file
// @Failure      500  {object}  errors.Error
// @Router /browse/thumbnail [get]
func GetThumbnailHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("name")

	log.Info().
		Str("name", item).
		Msg("Thumbnail")

	m, err := meta.Read(r.Context(), handler.EntClient(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(m.Thumbnail)
}
