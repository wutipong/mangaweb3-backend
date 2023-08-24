package tag

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/tag"
)

type setTagFavoriteRequest struct {
	Tag      string `json:"tag"`
	Favorite bool   `json:"favorite"`
}

type setTagFavoriteResponse struct {
	Request  setTagFavoriteRequest `json:"request"`
	Favorite bool                  `json:"favorite"`
}

func SetFavoriteHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := setTagFavoriteRequest{}
	if err := handler.ParseInput(r.Body, req); err != nil {
		handler.WriteResponse(w, err)
	}

	log.Info().Interface("request", req).Msg("Set favorite tag.")

	m, err := tag.Read(r.Context(), req.Tag)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if req.Favorite != m.Favorite {
		m.Favorite = req.Favorite
		tag.Write(r.Context(), m)
	}

	response := setTagFavoriteResponse{
		Request:  req,
		Favorite: m.Favorite,
	}

	handler.WriteResponse(w, response)
}
