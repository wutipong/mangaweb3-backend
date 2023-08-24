package browse

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/meta"
)

func ThumbnailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	item := r.URL.Query().Get("name")

	log.Info().
		Str("name", item).
		Msg("Thumbnail")

	m, err := meta.Read(r.Context(), item)
	if errors.Is(err, sql.ErrNoRows) {
		m, _ = meta.NewItem(item)
		err = meta.Write(r.Context(), m)
		if err != nil {
			handler.WriteResponse(w, err)
			return
		}

	} else if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(m.Thumbnail)
}
