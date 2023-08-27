package view

import (
	_ "image/png"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/meta"
)

type updateCoverRequest struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type updateCoverResponse struct {
	Request updateCoverRequest `json:"request"`
	Success bool               `json:"success"`
}

const (
	PathUpdateCover = "/view/update_cover"
)

// @accept json
// @Param request body view.updateCoverRequest true "request"
// @Success      200  {object}  view.updateCoverResponse
// @Failure      500  {object}  errors.Error
// @Router /view/update_cover [post]
// UpdateCover a handler to update the cover to specific image
func UpdateCover(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := updateCoverRequest{}
	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	log.Info().
		Interface("request", req).
		Msg("Update cover.")

	m, err := meta.Read(r.Context(), req.Name)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	entryIndex := m.FileIndices[req.Index]

	err = meta.GenerateThumbnail(m, entryIndex)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	_, err = meta.Write(r.Context(), m)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	handler.WriteResponse(w, updateCoverResponse{
		Request: req,
		Success: true,
	})
}
