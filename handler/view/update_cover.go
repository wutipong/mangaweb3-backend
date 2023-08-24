package view

import (
	"encoding/json"
	_ "image/png"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"go.uber.org/zap"
)

type updateCoverRequest struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type updateCoverResponse struct {
	Request updateCoverRequest `json:"request"`
	Success bool               `json:"success"`
}

// UpdateCover a handler to update the cover to specific image
func UpdateCover(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := updateCoverRequest{}
	if reqBody, err := io.ReadAll(r.Body); err != nil {
		handler.WriteResponse(w, err)
	} else {
		json.Unmarshal(reqBody, &req)
	}

	item := req.Name
	index := req.Index

	log.Get().Info("Update Cover", zap.String("item_name", item), zap.Int("index", index))

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	entryIndex := m.FileIndices[index]

	err = m.GenerateThumbnail(entryIndex)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	err = meta.Write(r.Context(), m)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	handler.WriteResponse(w, updateCoverResponse{
		Request: req,
		Success: true,
	})
}
