package tag

import (
	"hash/fnv"
	"net/http"
	"sort"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/tag"
)

const (
	ItemPerPage = 40
)

type PageData struct {
	Version string
	Tags    []ItemData
}

type ItemData struct {
	ID       uint64
	Name     string
	Favorite bool
}

func createItems(allTags []tag.Tag, favoriteOnly bool) []ItemData {
	allItems := make([]ItemData, len(allTags))

	for i, t := range allTags {
		isAdding := true
		if favoriteOnly {
			isAdding = t.Favorite
		}

		if isAdding {
			hash := fnv.New64()
			hash.Write([]byte(t.Name))
			id := hash.Sum64()

			allItems[i] = ItemData{
				ID:       id,
				Name:     t.Name,
				Favorite: t.Favorite,
			}
		}
	}

	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].Name < allItems[j].Name
	})
	return allItems
}

func TagListHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()

	log.Info().Msg("Tag list")

	favOnly := false
	if f, e := strconv.ParseBool(query.Get("favorite")); e == nil {
		favOnly = f
	}

	allTags, err := tag.ReadAll(r.Context())
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	tagData := createItems(allTags, favOnly)

	data := PageData{
		Version: handler.CreateVersionString(),
		Tags:    tagData,
	}

	handler.WriteResponse(w, data)
}
