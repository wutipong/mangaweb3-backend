package view

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog/log"
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
	BrowseURL string      `json:"browse_url"`
	Favorite  bool        `json:"favorite"`
	Indices   []int       `json:"indices"`
	Tags      []string    `json:"tags"`
}

const (
	PathView = "/view"
)

// @accept json
// @Param request body view.viewRequest true "request"
// @Success      200  {object}  view.viewResponse
// @Failure      500  {object}  errors.Error
// @Router /view [post]
func Handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := viewRequest{}

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

	id := m.ID

	if !m.Read {
		m.Read = true
		meta.Write(r.Context(), m)
	}

	browseUrl := r.Referer()
	if browseUrl == "" {
		browseUrl = handler.CreateBrowseURL(strconv.FormatInt(int64(id), 16))
	} else {
		if u, e := url.Parse(browseUrl); e == nil {
			u.Fragment = strconv.FormatInt(int64(id), 16)
			browseUrl = u.String()
		}
	}

	log.Info().
		Interface("request", req).
		Msg("View Item")

	data := viewResponse{
		Request:   req,
		Name:      item,
		Version:   handler.CreateVersionString(),
		BrowseURL: browseUrl,
		Favorite:  m.Favorite,
		Tags:      m.Tags,
		Indices:   m.FileIndices,
	}

	handler.WriteResponse(w, data)
}
