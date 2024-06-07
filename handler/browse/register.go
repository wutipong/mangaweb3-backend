package browse

import (
	"goji.io"
	"goji.io/pat"
)

func Register(mux *goji.Mux) {
	mux.HandleFunc(pat.Get(PathRecreateThumbnails), RecreateThumbnailHandler)
	mux.HandleFunc(pat.Get(PathRescanLibrary), RescanLibraryHandler)
	mux.HandleFunc(pat.Get(PathThumbnail), GetThumbnailHandler)
	mux.HandleFunc(pat.Post(PathBrowse), BrowseHandler)
	mux.HandleFunc(pat.Post(PathHistory), HistoryHandler)
}
