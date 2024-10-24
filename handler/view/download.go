package view

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/data"
	"github.com/wutipong/mangaweb3-backend/database"
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

	client := database.CreateEntClient()
	defer client.Close()

	m, err := meta.Read(r.Context(), client, item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	reader, err := data.GetObject(r.Context(), m.Name)

	bytes, err := io.ReadAll(reader)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(m.Name)))
	w.WriteHeader(http.StatusOK)

	w.Write(bytes)
}
