package browse

import (
	"net/http"

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
	ItemPerPage  int            `json:"item_per_page" default:"30"`
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
	// ID of the item
	ID int `json:"id,omitempty"`
	// Name of the item
	Name string `json:"name,omitempty"`
	// Favorite this item is a favorite
	Favorite bool `json:"favorite,omitempty"`
	// Read this item has been read before.
	Read bool `json:"read,omitempty"`
	// PageCount the number of pages.
	PageCount int `json:"page_count,omitempty"`
	// TagFavorite this item contains favorite tags
	TagFavorite bool `json:"tag_favorite,omitempty"`
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
func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	req := createDefaultBrowseRequest()

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
	}

	allMeta, err := meta.ReadPage(r.Context(),
		handler.EntClient(),
		meta.QueryParams{
			SearchName:   req.Search,
			FavoriteOnly: req.FavoriteOnly,
			SearchTag:    req.Tag,
			SortBy:       req.Sort,
			SortOrder:    req.Order,
			Page:         req.Page,
			ItemPerPage:  req.ItemPerPage,
		})
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	items := make([]browseItem, len(allMeta))
	for i, m := range allMeta {
		items[i] = browseItem{
			ID:        m.ID,
			Name:      m.Name,
			Favorite:  m.Favorite,
			Read:      m.Read,
			PageCount: len(m.FileIndices),
		}

		tags, err := m.QueryTags().All(r.Context())
		if err != nil {
			handler.WriteResponse(w, err)
			return
		}

		for _, t := range tags {
			if t.Favorite {
				items[i].TagFavorite = true
				break
			}
		}
	}

	count, err := meta.Count(r.Context(),
		handler.EntClient(),
		meta.QueryParams{
			SearchName:   req.Search,
			FavoriteOnly: req.FavoriteOnly,
			SearchTag:    req.Tag,
			SortBy:       req.Sort,
			SortOrder:    req.Order,
			Page:         0,
			ItemPerPage:  0,
		})
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
		tagObj, err := tag.Read(r.Context(), handler.EntClient(), req.Tag)
		if err != nil {
			handler.WriteResponse(w, err)
			return
		}

		data.TagFavorite = tagObj.Favorite
	}

	handler.WriteResponse(w, data)
}
