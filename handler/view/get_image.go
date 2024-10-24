package view

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"

	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/data"
	"github.com/wutipong/mangaweb3-backend/database"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"

	_ "golang.org/x/image/webp"
)

const (
	PathGetImage = "/view/get_image"
)

// @Param name query string true "name of the item"
// @Param width query int false "width"
// @Param height query int false "height"
// @Param i query int true "index"
// @Success      200  {body}  file
// @Failure      500  {object}  errors.Error
// @Router /view/get_image [get]
// GetImage returns an image with specific width/height while retains aspect ratio.
func GetImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()
	item := query.Get("name")

	var width, height int64 = 0, 0
	if w, e := strconv.ParseInt(query.Get("width"), 10, 64); e == nil {
		width = w
		height = width
	}

	if h, e := strconv.ParseInt(query.Get("height"), 10, 64); e == nil {
		height = h
	}

	var index = 0
	if i, err := strconv.Atoi(query.Get("i")); err == nil {
		index = i
	}

	log.Info().
		Interface("request", query).
		Msg("Get image")

	client := database.CreateEntClient()
	defer client.Close()

	m, err := meta.Read(r.Context(), client, item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}
	data, f, err := OpenZipEntry(r.Context(), m, index)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if width == 0 && height == 0 {
		var contentType string
		switch filepath.Ext(strings.ToLower(f)) {
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		case ".webp":
			contentType = "image/webp"
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
		w.Header().Set("Content-Type", contentType)

		return
	}

	reader := bytes.NewBuffer(data)

	img, err := imaging.Decode(reader, imaging.AutoOrientation(true))
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if img.Bounds().Dx() > int(width) || img.Bounds().Dy() > int(height) {
		resized := imaging.Fit(img, int(width), int(height), imaging.MitchellNetravali)
		img = resized
	}

	w.WriteHeader(http.StatusOK)
	err = imaging.Encode(w, img, imaging.JPEG)

	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
}

func OpenZipEntry(ctx context.Context, m *ent.Meta, index int) (content []byte, filename string, err error) {
	if len(m.FileIndices) == 0 {
		err = fmt.Errorf("image file not found")
	}

	children, err := data.ListObject(ctx, m.Name)
	if err != nil {
		return
	}

	filename = children[m.FileIndices[index]]
	reader, err := data.GetObject(ctx, path.Join(m.Name, filename))
	if err != nil {
		return
	}
	defer reader.Close()
	if content, err = io.ReadAll(reader); err != nil {
		content = nil
		return
	}
	return
}
