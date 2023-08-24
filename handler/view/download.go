package view

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/meta"
)

const (
	PathDownload = "/view/download"
)

// @Param name query string true "name of the file"
// @Success      200  {body}  file
// @Failure      500  {object}  errors.Error
// @Router /view/download [get]
func Download(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	item := r.URL.Query().Get("name")

	log.Info().Str("name", item).Msg("Download")

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	reader, err := m.Open()
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}
	defer reader.Close()

	bytes, err := io.ReadAll(reader)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
}
