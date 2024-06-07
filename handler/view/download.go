package view

import (
	"io"
	"net/http"
	"strconv"

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
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("name")

	log.Info().Str("name", item).Msg("Download")

	m, err := meta.Read(r.Context(), handler.EntClient(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	reader, err := meta.Open(m)
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
