package browse

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/tag"
)

const (
	ItemPerPage = 30
)

type browseRequest struct {
	Tag          string         `json:"tag"`
	FavoriteOnly bool           `json:"favorite_only"`
	Page         int            `json:"page"`
	Search       string         `json:"search"`
	Sort         meta.SortField `json:"sort"`
	Order        meta.SortOrder `json:"order"`
}

type browseData struct {
	Request           browseRequest `json:"request"`
	Title             string
	Version           string
	FavoriteOnly      bool
	SortBy            string
	SortOrder         string
	Tag               string
	TagFavorite       bool
	SetTagFavoriteURL string
	BrowseURL         string
	TagListURL        string
	SearchText        string
	RescanURL         string
	Items             []item
	Pages             []pageItem
}

type item struct {
	ID         uint64
	Name       string
	CreateTime time.Time
	Favorite   bool
	IsRead     bool
}

type pageItem struct {
	Content         string
	LinkURL         string
	IsActive        bool
	IsEnabled       bool
	IsHiddenOnSmall bool
}

func createItems(allMeta []meta.Meta) (allItems []item, err error) {
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
			IsRead:     m.IsRead,
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
	}
}

func Handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := createDefaultBrowseRequest()

	if err := handler.ParseInput(r.Body, req); err != nil {
		handler.WriteResponse(w, err)
	}

	tagStr := req.Tag
	favOnly := req.FavoriteOnly
	page := req.Page
	search := req.Search
	searchCriteria := make([]meta.SearchCriteria, 0)
	if search != "" {
		searchCriteria = append(searchCriteria, meta.SearchCriteria{
			Field: meta.SearchFieldName,
			Value: search,
		})
	}

	if favOnly {
		searchCriteria = append(searchCriteria, meta.SearchCriteria{
			Field: meta.SearchFieldFavorite,
			Value: true,
		})
	}

	if tagStr != "" {
		searchCriteria = append(searchCriteria, meta.SearchCriteria{
			Field: meta.SearchFieldTag,
			Value: tagStr,
		})
	}

	sort := req.Sort
	order := req.Order

	allMeta, err := meta.Search(r.Context(), searchCriteria, sort, order, ItemPerPage, page)
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

	pageCount := int(count / ItemPerPage)
	if count%ItemPerPage > 0 {
		pageCount++
	}

	if page > pageCount || page < 0 {
		page = 0
	}

	log.Info().
		Interface("request", req).
		Msg("Browse")

	data := browseData{
		Request:      req,
		Title:        "Browse - All items",
		Version:      handler.CreateVersionString(),
		FavoriteOnly: favOnly,
		SortBy:       string(sort),
		SortOrder:    string(order),
		Items:        items,
		Pages:        createPageItems(page, pageCount, r.URL),
		BrowseURL:    handler.CreateBrowseURL(""),
		TagListURL:   handler.CreateTagListURL(),
		SearchText:   search,
		RescanURL:    handler.CreateRescanURL(),
	}

	if tagStr != "" {
		data.Title = fmt.Sprintf("Browse - %s", tagStr)
		data.Tag = tagStr

		tagObj, err := tag.Read(r.Context(), tagStr)
		if err != nil {
			handler.WriteResponse(w, err)
			return
		}

		data.TagFavorite = tagObj.Favorite
		data.SetTagFavoriteURL = handler.CreateSetTagFavoriteURL(tagStr)
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
