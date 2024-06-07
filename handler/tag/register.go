package tag

import (
	"goji.io"
	"goji.io/pat"
)

func Register(mux *goji.Mux) {
	mux.HandleFunc(pat.Get(PathThumbnail), ThumbnailHandler)
	mux.HandleFunc(pat.Get(PathRecreateThumbnails), RecreateThumbnailHandler)
	mux.HandleFunc(pat.Post(PathList), ListHandler)
	mux.HandleFunc(pat.Post(PathSetFavorite), SetFavoriteHandler)
}
