package view

import (
	"goji.io"
	"goji.io/pat"
)

func Register(mux *goji.Mux) {
	mux.HandleFunc(pat.Get(PathDownload), DownloadHandler)
	mux.HandleFunc(pat.Get(PathGetImage), GetImageHandler)
	mux.HandleFunc(pat.Post(PathFavorite), SetFavoriteHandler)
	mux.HandleFunc(pat.Post(PathUpdateCover), UpdateCoverHandler)
	mux.HandleFunc(pat.Post(PathView), ViewHandler)
	mux.HandleFunc(pat.Post(PathPopulateTags), PopulateTagsHandler)
	mux.HandleFunc(pat.Post(PathFixMeta), FixMetaHandler)
}
