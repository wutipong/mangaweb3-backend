package view

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/tag"
)

const (
	maxPageWidth  = 1600
	maxPageHeight = 1600
)

type viewRequest struct {
	Path string `json:"path"`
}

type viewData struct {
	Request          viewRequest `json:"request"`
	Name             string
	Title            string
	Version          string
	BrowseURL        string
	Favorite         bool
	ImageURLs        []string
	UpdateCoverURLs  []string
	DownloadPageURLs []string
	Tags             []tagData
	DownloadURL      string
	SetFavoriteURL   string
}

type tagData struct {
	Name string
	URL  string
}

func Handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := viewRequest{}

	if reqBody, err := io.ReadAll(r.Body); err != nil {
		handler.WriteResponse(w, err)
		return
	} else {
		json.Unmarshal(reqBody, &req)
	}

	item := req.Path

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	pages, err := ListPages(m)
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

	tags := make([]tagData, 0)

	for _, tagStr := range m.Tags {
		t, err := tag.Read(r.Context(), tagStr)
		if err != nil {
			log.Get().Sugar().Fatal(err)
			handler.WriteResponse(w, err)
			return
		}

		if !t.Hidden {
			tags = append(tags, tagData{
				Name: t.Name,
				URL:  handler.CreateBrowseTagURL(t.Name),
			})
		}
	}

	log.Get().Info("View Item", zap.String("item_name", item))

	data := viewData{
		Request:          req,
		Name:             item,
		Title:            fmt.Sprintf("Manga - Viewing [%s]", item),
		Version:          handler.CreateVersionString(),
		BrowseURL:        browseUrl,
		ImageURLs:        createImageURLs(item, pages),
		UpdateCoverURLs:  createUpdateCoverURLs(item, pages),
		DownloadPageURLs: createDownloadImageURLs(item, pages),
		Favorite:         m.Favorite,
		Tags:             tags,
		DownloadURL:      handler.CreateDownloadURL(item),
		SetFavoriteURL:   handler.CreateSetFavoriteURL(item),
	}

	handler.WriteResponse(w, data)
}

func createDownloadImageURLs(file string, pages []Page) []string {
	output := make([]string, len(pages))
	for i, p := range pages {
		output[i] = handler.CreateGetImageURL(file, p.Index)
	}
	return output
}

func createImageURLs(file string, pages []Page) []string {
	output := make([]string, len(pages))
	for i, p := range pages {
		output[i] = handler.CreateGetImageWithSizeURL(file, p.Index, maxPageWidth, maxPageHeight)
	}
	return output
}

func createUpdateCoverURLs(file string, pages []Page) []string {
	output := make([]string, len(pages))
	for i, p := range pages {
		output[i] = handler.CreateUpdateCoverURL(file, p.Index)
	}
	return output
}
