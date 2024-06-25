package maintenance

import "github.com/julienschmidt/httprouter"

func Register(router *httprouter.Router) {
	router.GET(PathUpdateLibrary, UpdateLibraryHandler)
}
