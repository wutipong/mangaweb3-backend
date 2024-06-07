package view

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
)

type setFavoriteRequest struct {
	Name     string `json:"name"`
	Favorite bool   `json:"favorite"`
}

type setFavoriteResponse struct {
	Request  setFavoriteRequest `json:"request"`
	Favorite bool               `json:"favorite"`
}

const (
	PathFavorite = "/view/set_favorite"
)

// @accept json
// @Param request body view.setFavoriteRequest true "request"
// @Success      200  {object}  view.setFavoriteResponse
// @Failure      500  {object}  errors.Error
// @Router /view/set_favorite [post]
func SetFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	req := setFavoriteRequest{}
	var reqBody []byte

	if read, err := io.ReadAll(r.Body); err != nil {
		handler.WriteResponse(w, err)
		return
	} else {
		reqBody = read
	}

	if err := json.Unmarshal(reqBody, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	item := req.Name

	log.Info().
		Interface("request", req).
		Msg("Set Favorite Item.")

	m, err := meta.Read(r.Context(), handler.EntClient(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if req.Favorite != m.Favorite {
		m.Favorite = req.Favorite
		meta.Write(r.Context(), handler.EntClient(), m)
	}

	response := setFavoriteResponse{
		Request:  req,
		Favorite: m.Favorite,
	}

	handler.WriteResponse(w, response)
}
