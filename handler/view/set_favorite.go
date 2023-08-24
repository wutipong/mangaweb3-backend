package view

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"go.uber.org/zap"
)

type setFavoriteRequest struct {
	Name     string `json:"name"`
	Favorite bool   `json:"favorite"`
}

type setFavoriteResponse struct {
	Request  setFavoriteRequest `json:"request"`
	Favorite bool               `json:"favorite"`
}

func SetFavoriteHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

	log.Get().Info("Set Favorite Item", zap.String("item_name", item))

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if req.Favorite != m.Favorite {
		m.Favorite = req.Favorite
		meta.Write(r.Context(), m)
	}

	response := setFavoriteResponse{
		Request:  req,
		Favorite: m.Favorite,
	}

	handler.WriteResponse(w, response)
}
