package browse

import (
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent/history"
	"github.com/wutipong/mangaweb3-backend/handler"
)

const (
	PathHistory = "/history"
)

type historyRequest struct {
	Page        int `json:"page"`
	ItemPerPage int `json:"item_per_page" default:"30"`
}

type historyResponse struct {
	Request   historyRequest `json:"request"`
	TotalPage int            `json:"total_page"`
	Items     []historyItem  `json:"items"`
}

type historyItem struct {
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
	// AccessTime the time the item is accessed.
	AccessTime time.Time `json:"access_time,omitempty"`
	// TagFavorite this item contains favorite tags
	TagFavorite bool `json:"tag_favorite,omitempty"`
}

func createDefaultHistoryRequest() historyRequest {
	return historyRequest{
		Page:        0,
		ItemPerPage: 30,
	}
}

// @accept json
// @Param request body browse.historyRequest false "request"
// @Success      200  {object}  browse.historyResponse
// @Failure      500  {object}  errors.Error
// @Router /history [post]
func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	req := createDefaultHistoryRequest()

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
	}

	histories, err := handler.EntClient().History.Query().
		Order(history.ByCreateTime(sql.OrderDesc())).
		Limit(req.ItemPerPage).
		Offset(req.ItemPerPage * req.Page).All(r.Context())

	if err != nil {
		handler.WriteResponse(w, err)
	}

	items := make([]historyItem, len(histories))

	for i, h := range histories {
		m, err := h.QueryItem().Only(r.Context())

		if err != nil {
			handler.WriteResponse(w, err)
			return
		}

		items[i] = historyItem{
			ID:         m.ID,
			Name:       m.Name,
			Favorite:   m.Favorite,
			Read:       m.Read,
			PageCount:  len(m.FileIndices),
			AccessTime: h.CreateTime,
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

	count, err := handler.EntClient().History.Query().Count(r.Context())

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

	data := historyResponse{
		Request:   req,
		Items:     items,
		TotalPage: pageCount,
	}

	handler.WriteResponse(w, data)
}
