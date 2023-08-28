package tag

import "github.com/julienschmidt/httprouter"

func Register(router *httprouter.Router) {
	router.GET(PathThumbnail, ThumbnailHandler)
	router.GET(PathRecreateThumbnails, RecreateThumbnailHandler)
	router.POST(PathList, ListHandler)
	router.POST(PathSetFavorite, SetFavoriteHandler)
}
