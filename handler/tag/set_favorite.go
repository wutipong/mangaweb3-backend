package tag

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/tag"
)

type setFavoriteRequest struct {
	Tag      string `json:"tag"`
	Favorite bool   `json:"favorite"`
}

type setFavoriteResponse struct {
	Request  setFavoriteRequest `json:"request"`
	Favorite bool               `json:"favorite"`
}

const (
	PathSetFavorite = "/tag/set_favorite"
)

// @accept json
// @Param request body tag.setFavoriteRequest true "request"
// @Success      200  {object}  tag.setFavoriteResponse
// @Failure      500  {object}  errors.Error
// @Router /tag/set_favorite [post]
func SetFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	req := setFavoriteRequest{}
	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
	}

	log.Info().Interface("request", req).Msg("Set favorite tag.")

	m, err := tag.Read(r.Context(), handler.EntClient(), req.Tag)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if req.Favorite != m.Favorite {
		m.Favorite = req.Favorite
		tag.Write(r.Context(), handler.EntClient(), m)
	}

	response := setFavoriteResponse{
		Request:  req,
		Favorite: m.Favorite,
	}

	handler.WriteResponse(w, response)
}
