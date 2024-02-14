package tag

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/tag"
)

type listRequest struct {
	FavoriteOnly bool   `json:"favorite_only"`
	Search       string `json:"search"`
	Page         int    `json:"page"`
	ItemPerPage  int    `json:"item_per_page"`
}

type listResponse struct {
	Request   listRequest `json:"request"`
	Tags      []*ent.Tag  `json:"tags"`
	TotalPage int         `json:"total_page"`
}

const (
	PathList = "/tag/list"
)

const (
	DefaultItemPerPage = 30
)

// @accept json
// @Param request body tag.listRequest true "request"
// @Success      200  {object}  tag.listResponse
// @Failure      500  {object}  errors.Error
// @Router /tag/list [post]
func ListHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := listRequest{
		FavoriteOnly: false,
		Search:       "",
		Page:         0,
		ItemPerPage:  DefaultItemPerPage,
	}

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	log.Info().Interface("request", req).Msg("Tag list")

	allTags, err := tag.ReadPage(r.Context(), handler.EntClient(), req.FavoriteOnly, req.Search, req.Page, req.ItemPerPage)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	total, err := tag.Count(r.Context(), handler.EntClient(), req.FavoriteOnly, req.Search)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	data := listResponse{
		Request:   req,
		Tags:      allTags,
		TotalPage: (total / req.ItemPerPage) + 1,
	}

	handler.WriteResponse(w, data)
}
