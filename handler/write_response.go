package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wutipong/mangaweb3-backend/errors"
)

func WriteResponse(w http.ResponseWriter, v any) {
	if err, ok := v.(error); ok {
		w.WriteHeader(http.StatusInternalServerError)
		if _, ok := err.(errors.Error); !ok {
			v = errors.ErrUnknown.Wrap(err)
		}

	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, _ := json.Marshal(v)
	w.Write(b)
}
