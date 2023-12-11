package browse

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/tag"
)

const (
	PathBrowse = "/browse"
)

type browseRequest struct {
	Tag          string         `json:"tag"`
	FavoriteOnly bool           `json:"favorite_only"`
	Page         int            `json:"page"`
	ItemPerPage  int            `json:"item_per_page"`
	Search       string         `json:"search"`
	Sort         meta.SortField `json:"sort"`
	Order        meta.SortOrder `json:"order"`
}

type browseResponse struct {
	Request     browseRequest `json:"request"`
	TagFavorite bool          `json:"tag_favorite"`
	TotalPage   int           `json:"total_page"`
	Items       []browseItem  `json:"items"`
}

type browseItem struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Favorite holds the value of the "favorite" field.
	Favorite bool `json:"favorite,omitempty"`
	// Read holds the value of the "read" field.
	Read bool `json:"read,omitempty"`
}

func createDefaultBrowseRequest() browseRequest {
	return browseRequest{
		Tag:          "",
		FavoriteOnly: false,
		Page:         0,
		Search:       "",
		Sort:         meta.SortFieldCreateTime,
		Order:        meta.SortOrderDescending,
		ItemPerPage:  30,
	}
}

// @accept json
// @Param request body browse.browseRequest false "request"
// @Success      200  {object}  browse.browseResponse
// @Failure      500  {object}  errors.Error
// @Router /browse [post]
func Handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := createDefaultBrowseRequest()

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
	}

	allMeta, err := meta.SearchItems(r.Context(),
		req.Search,
		req.FavoriteOnly,
		req.Tag,
		req.Sort,
		req.Order,
		req.Page,
		req.ItemPerPage)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	items := make([]browseItem, len(allMeta))
	for i, m := range allMeta {
		items[i] = browseItem{
			ID:       m.ID,
			Name:     m.Name,
			Favorite: m.Favorite,
			Read:     m.Read,
		}
	}

	count, err := meta.CountItems(r.Context(),
		req.Search,
		req.FavoriteOnly,
		req.Tag,
		req.Sort,
		req.Order)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	pageCount := int(count) / req.ItemPerPage
	if int(count)%req.ItemPerPage > 0 {
		pageCount++
	}

	if req.Page > pageCount || req.Page < 0 {
		req.Page = 0
	}

	log.Info().
		Interface("request", req).
		Msg("Browse")

	data := browseResponse{
		Request:   req,
		Items:     items,
		TotalPage: pageCount,
	}

	if req.Tag != "" {
		tagObj, err := tag.Read(r.Context(), req.Tag)
		if err != nil {
			handler.WriteResponse(w, err)
			return
		}

		data.TagFavorite = tagObj.Favorite
	}

	handler.WriteResponse(w, data)
}
