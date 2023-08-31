package view

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
)

type populateTagsRequest struct {
	Name string `json:"name"`
}

type populateTagsResponse struct {
	Request  populateTagsRequest `json:"request"`
	Name     string              `json:"name"`
	Version  string              `json:"version"`
	Favorite bool                `json:"favorite"`
	Indices  []int               `json:"indices"`
	Tags     []*ent.Tag          `json:"tags"`
}

const (
	PathPopulateTags = "/view/populate_tags"
)

// @accept json
// @Param request body view.populateTagsRequest true "request"
// @Success      200  {object}  view.populateTagsResponse
// @Failure      500  {object}  errors.Error
// @Router /view/populate_tags [post]
func PopulateTagsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := populateTagsRequest{}

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	item := req.Name

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	log.Info().
		Interface("request", req).
		Msg("View Item")

	m, err = meta.PopulateTags(r.Context(), m)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	tags, err := m.QueryTags().All(r.Context())
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	data := populateTagsResponse{
		Request:  req,
		Name:     item,
		Version:  handler.CreateVersionString(),
		Favorite: m.Favorite,
		Tags:     tags,
		Indices:  m.FileIndices,
	}

	handler.WriteResponse(w, data)
}
