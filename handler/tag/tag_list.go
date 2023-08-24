package tag

import (
	"hash/fnv"
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/handler"

	"github.com/wutipong/mangaweb3-backend/tag"
)

const (
	ItemPerPage = 40
)

type listRequest struct {
	FavoriteOnly bool `json:"favorite_only"`
}

type listResponse struct {
	Request listRequest `json:"request"`
	Tags    []tagData   `json:"tags"`
}

type tagData struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Favorite bool   `json:"favorite"`
}

func createItems(allTags []tag.Tag, favoriteOnly bool) []tagData {
	allItems := make([]tagData, len(allTags))

	for i, t := range allTags {
		isAdding := true
		if favoriteOnly {
			isAdding = t.Favorite
		}

		if isAdding {
			hash := fnv.New64()
			hash.Write([]byte(t.Name))
			id := hash.Sum64()

			allItems[i] = tagData{
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

// @accept json
// @Param request body tag.listRequest true "request"
// @Success      200  {object}  tag.listResponse
// @Failure      500  {object}  errors.Error
// @Router /tag/list [post]
func ListHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := listRequest{
		FavoriteOnly: false,
	}

	if err := handler.ParseInput(r.Body, &req); err != nil {
		handler.WriteResponse(w, err)
		return
	}

	log.Info().Interface("request", req).Msg("Tag list")

	allTags, err := tag.ReadAll(r.Context())
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	tagData := createItems(allTags, req.FavoriteOnly)

	data := listResponse{
		Tags: tagData,
	}

	handler.WriteResponse(w, data)
}
