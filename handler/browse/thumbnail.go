package browse

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"go.uber.org/zap"
)

type thumbnailRequest struct {
	Path string `json:"path"`
}

func ThumbnailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := thumbnailRequest{}
	if reqBody, err := io.ReadAll(r.Body); err != nil {
		handler.WriteResponse(w, err)
	} else {
		json.Unmarshal(reqBody, &req)
	}

	item := req.Path
	item = filepath.FromSlash(item)

	log.Get().Info("Item Thumbnail", zap.String("item_name", item))

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
