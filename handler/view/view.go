package view

import (
	"hash/fnv"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/meta"
)

const (
	maxPageWidth  = 1600
	maxPageHeight = 1600
)

type viewRequest struct {
	Path string `json:"path"`
}

type viewData struct {
	Request   viewRequest `json:"request"`
	Name      string      `json:"name"`
	Version   string      `json:"version"`
	BrowseURL string      `json:"browse_url"`
	Favorite  bool        `json:"favorite"`
	Indices   []int       `json:"indices"`
	Tags      []string    `json:"tags"`
}

func Handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := viewRequest{}

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	item := req.Path

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	hash := fnv.New64()
	hash.Write([]byte(item))
	id := hash.Sum64()

	if !m.IsRead {
		m.IsRead = true
		meta.Write(r.Context(), m)
	}

	browseUrl := r.Referer()
	if browseUrl == "" {
		browseUrl = handler.CreateBrowseURL(strconv.FormatUint(id, 16))
	} else {
		if u, e := url.Parse(browseUrl); e == nil {
			u.Fragment = strconv.FormatUint(id, 10)
			browseUrl = u.String()
		}
	}

	log.Info().
		Interface("request", req).
		Msg("View Item")

	data := viewData{
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
