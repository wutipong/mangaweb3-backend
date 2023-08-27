package browse

import (
	"hash/fnv"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
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
	Items       []item        `json:"items"`
	Pages       []pageItem    `json:"pages"`
}

type item struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	Favorite   bool      `json:"favorite"`
	IsRead     bool      `json:"is_read"`
}

type pageItem struct {
	Content         string `json:"content"`
	LinkURL         string `json:"link_url"`
	IsActive        bool   `json:"is_active"`
	IsEnabled       bool   `json:"is_enabled"`
	IsHiddenOnSmall bool   `json:"is_hidden_on_small"`
}

func createItems(allMeta []*ent.Meta) (allItems []item, err error) {
	allItems = make([]item, len(allMeta))

	for i, m := range allMeta {
		hash := fnv.New64()
		hash.Write([]byte(m.Name))
		id := hash.Sum64()

		allItems[i] = item{
			ID:         id,
			Name:       m.Name,
			CreateTime: m.CreateTime,
			Favorite:   m.Favorite,
			IsRead:     m.Read,
		}
	}
	return
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

	search := req.Search
	searchCriteria := make([]meta.SearchCriteria, 0)
	if search != "" {
		searchCriteria = append(searchCriteria, meta.SearchCriteria{
			Field: meta.SearchFieldName,
			Value: search,
		})
	}

	if req.FavoriteOnly {
		searchCriteria = append(searchCriteria, meta.SearchCriteria{
			Field: meta.SearchFieldFavorite,
			Value: true,
		})
	}

	if req.Tag != "" {
		searchCriteria = append(searchCriteria, meta.SearchCriteria{
			Field: meta.SearchFieldTag,
			Value: req.Tag,
		})
	}

	allMeta, err := meta.Search(r.Context(), searchCriteria, req.Sort, req.Order, req.ItemPerPage, req.Page)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	items, err := createItems(allMeta)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	count, err := meta.Count(r.Context(), searchCriteria)
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
		Pages:     createPageItems(req.Page, pageCount, r.URL),
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

func createPageItems(current int, count int, baseUrl *url.URL) []pageItem {
	const (
		First    = "First"
		Previous = "Previous"
		Next     = "Next"
		Last     = "Last"

		DisplayPageCount     = 6
		HalfDisplayPageCount = DisplayPageCount / 2
	)

	firstPage := 0
	lastPage := count - 1
	previousPage := current - 1
	nextPage := current + 1

	changePageParam := func(baseUrl *url.URL, page int) *url.URL {
		query := baseUrl.Query()

		if query.Has("page") {
			query.Set("page", strconv.Itoa(page))
		} else {
			query.Add("page", strconv.Itoa(page))
		}

		baseUrl.RawQuery = query.Encode()
		return baseUrl
	}

	output := make([]pageItem, 0)
	output = append(output, pageItem{
		Content:         First,
		LinkURL:         changePageParam(baseUrl, firstPage).String(),
		IsActive:        false,
		IsEnabled:       true,
		IsHiddenOnSmall: false,
	})

	enablePrevious := previousPage >= firstPage
	output = append(output, pageItem{
		Content:         Previous,
		LinkURL:         changePageParam(baseUrl, previousPage).String(),
		IsActive:        false,
		IsEnabled:       enablePrevious,
		IsHiddenOnSmall: false,
	})

	for i := current - HalfDisplayPageCount; i <= current+HalfDisplayPageCount; i++ {
		if i < firstPage {
			continue
		}
		if i > lastPage {
			continue
		}

		output = append(output, pageItem{
			Content:         strconv.Itoa(i),
			LinkURL:         changePageParam(baseUrl, i).String(),
			IsActive:        i == current,
			IsEnabled:       true,
			IsHiddenOnSmall: !(i == current),
		})
	}

	enableNext := nextPage < count
	output = append(output, pageItem{
		Content:         Next,
		LinkURL:         changePageParam(baseUrl, nextPage).String(),
		IsActive:        false,
		IsEnabled:       enableNext,
		IsHiddenOnSmall: false,
	})

	output = append(output, pageItem{
		Content:         Last,
		LinkURL:         changePageParam(baseUrl, lastPage).String(),
		IsActive:        false,
		IsEnabled:       true,
		IsHiddenOnSmall: false,
	})

	return output
}
