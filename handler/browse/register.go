package browse

import "github.com/julienschmidt/httprouter"

func Register(router *httprouter.Router) {
	router.GET(PathRecreateThumbnails, RecreateThumbnailHandler)
	router.GET(PathRescanLibrary, RescanLibraryHandler)
	router.GET(PathThumbnail, GetThumbnailHandler)
	router.POST(PathBrowse, Handler)
}
