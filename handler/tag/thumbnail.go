package tag

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/tag"
)

const (
	PathThumbnail = "/tag/thumbnail"
)

// @Param tag query string true "tag"
// @Success      200  {body}  file
// @Failure      500  {object}  errors.Error
// @Router /tag/thumbnail [get]
func ThumbnailHandler(w http.ResponseWriter, r *http.Request) {
	tagStr := r.URL.Query().Get("tag")

	log.Info().Str("tag", tagStr).Msg("Tag thumbnail image")

	m, err := tag.Read(r.Context(), handler.EntClient(), tagStr)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(m.Thumbnail)
	w.Header().Set("Content-Type", "image/jpeg")
}
