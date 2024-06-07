package view

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
)

type viewRequest struct {
	Name string `json:"name"`
}

type viewResponse struct {
	Request   viewRequest `json:"request"`
	Name      string      `json:"name"`
	Version   string      `json:"version"`
	Favorite  bool        `json:"favorite"`
	PageCount int         `json:"page_count"`
	Tags      []*ent.Tag  `json:"tags"`
}

const (
	PathView = "/view"
)

// @accept json
// @Param request body view.viewRequest true "request"
// @Success      200  {object}  view.viewResponse
// @Failure      500  {object}  errors.Error
// @Router /view [post]
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	req := viewRequest{}

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	item := req.Name

	m, err := meta.Read(r.Context(), handler.EntClient(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if !m.Read {
		m.Read = true
		meta.Write(r.Context(), handler.EntClient(), m)
	}

	log.Info().
		Interface("request", req).
		Msg("View Item")

	tags, err := m.QueryTags().All(r.Context())
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	data := viewResponse{
		Request:   req,
		Name:      item,
		Version:   handler.CreateVersionString(),
		Favorite:  m.Favorite,
		Tags:      tags,
		PageCount: len(m.FileIndices),
	}

	handler.EntClient().History.Create().
		SetItem(m).
		Save(r.Context())

	handler.WriteResponse(w, data)
}
