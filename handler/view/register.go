package view

import "github.com/julienschmidt/httprouter"

func Register(router *httprouter.Router) {
	router.GET(PathDownload, Download)
	router.GET(PathGetImage, GetImage)
	router.POST(PathFavorite, SetFavoriteHandler)
	router.POST(PathUpdateCover, UpdateCover)
	router.POST(PathView, Handler)
}
