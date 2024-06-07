package view

import (
	_ "image/png"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
)

type updateCoverRequest struct {
	Name    string           `json:"name"`
	Index   int              `json:"index"`
	Details meta.CropDetails `json:"crop_details"`
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
// UpdateCoverHandler a handler to update the cover to specific image
func UpdateCoverHandler(w http.ResponseWriter, r *http.Request) {
	req := updateCoverRequest{}
	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	log.Info().
		Interface("request", req).
		Msg("Update cover.")

	m, err := meta.Read(r.Context(), handler.EntClient(), req.Name)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	err = meta.GenerateThumbnail(m, req.Index, req.Details)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	err = meta.Write(r.Context(), handler.EntClient(), m)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	handler.WriteResponse(w, updateCoverResponse{
		Request: req,
		Success: true,
	})
}
