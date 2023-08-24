package tag

import (
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/tag"
)

func ThumbnailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tagStr := handler.ParseParam(params, "tag")
	tagStr = filepath.FromSlash(tagStr)

	log.Info().Str("tag", tagStr).Msg("Tag thumbnail image")

	m, err := tag.Read(r.Context(), tagStr)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(m.Thumbnail)
	w.Header().Set("Content-Type", "image/jpeg")
}
