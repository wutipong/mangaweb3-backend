package main

import (
	"context"
	"flag"
	"fmt"

	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/handler/browse"
	handlertag "github.com/wutipong/mangaweb3-backend/handler/tag"
	"github.com/wutipong/mangaweb3-backend/handler/view"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/scheduler"
	"github.com/wutipong/mangaweb3-backend/tag"
)

func setupFlag(flagName, defValue, variable, description string) *string {
	varValue := os.Getenv(variable)
	if varValue != "" {
		defValue = varValue
	}

	return flag.String(flagName, defValue, description)
}

var versionString string = "development"

const (
	pathBrowse        = "/browse"
	pathView          = "/view"
	pathGetImage      = "/get_image"
	pathUpdateCover   = "/update_cover"
	pathThumbnail     = "/thumbnail"
	pathFavorite      = "/favorite"
	pathDownload      = "/download"
	pathRescanLibrary = "/rescan_library"
	pathTagFavorite   = "/tag_favorite"
	pathTagList       = "/tag_list"
	pathTagThumb      = "/tag_thumb"
)

func main() {
	err := log.Init()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	err = godotenv.Load()
	if err != nil {
		log.Get().Info("Use .env file.")
	}

	address := setupFlag("address", ":8972", "MANGAWEB_ADDRESS", "The server address")
	dataPath := setupFlag("data", "./data", "MANGAWEB_DATA_PATH", "Manga source path")
	connectionStr := setupFlag(
		"database",
		"postgres://postgres:password@localhost:5432/manga",
		"MANGAWEB_DB",
		"Specify the database connection string",
	)

	flag.Parse()

	meta.BaseDirectory = *dataPath

	printBanner()
	log.Get().Sugar().Infof("MangaWeb version: %s", versionString)
	log.Get().Sugar().Infof("Data source Path: %s", *dataPath)
	log.Get().Sugar().Infof("Server starts at: %s", *address)

	router := httprouter.New()

	conn, err := pgxpool.New(context.Background(), *connectionStr)
	if err != nil {
		log.Get().Sugar().Fatal(err)

		return
	}

	defer conn.Close()

	tag.Init(conn)
	meta.Init(conn)

	scheduler.Init(scheduler.Options{})

	RegisterHandler(router)
	scheduler.Start()

	log.Get().Info("Server starts.")
	log.Get().Sugar().Fatal(http.ListenAndServe(*address, router))

	log.Get().Sugar().Info("shutting down the server")
	scheduler.Stop()
}

func RegisterHandler(router *httprouter.Router) {
	handler.Init(handler.Options{
		VersionString:     versionString,
		PathBrowse:        pathBrowse,
		PathView:          pathView,
		PathGetImage:      pathGetImage,
		PathUpdateCover:   pathUpdateCover,
		PathThumbnail:     pathThumbnail,
		PathFavorite:      pathFavorite,
		PathDownload:      pathDownload,
		PathRescanLibrary: pathRescanLibrary,
		PathTagFavorite:   pathTagFavorite,
		PathTagList:       pathTagList,
		PathTagThumbnail:  pathTagThumb,
	})
	// Routes
	router.POST(pathBrowse, browse.Handler)
	router.POST(pathView, view.Handler)
	router.GET(pathGetImage, handler.GetImage)
	router.POST(pathUpdateCover, view.UpdateCover)
	router.GET(pathThumbnail, browse.ThumbnailHandler)
	router.POST(pathFavorite, view.SetFavoriteHandler)
	router.GET(pathDownload, view.Download)
	router.GET(pathRescanLibrary, handler.RescanLibraryHandler)
	router.GET(pathTagFavorite, handlertag.SetFavoriteHandler)
	router.GET(pathTagList, handlertag.TagListHandler)
	router.GET(pathTagThumb, handlertag.ThumbnailHandler)
}

func printBanner() {
	fmt.Printf(
		`
	███╗░░░███╗░█████╗░███╗░░██╗░██████╗░░█████╗░░██╗░░░░░░░██╗███████╗██████╗░
	████╗░████║██╔══██╗████╗░██║██╔════╝░██╔══██╗░██║░░██╗░░██║██╔════╝██╔══██╗
	██╔████╔██║███████║██╔██╗██║██║░░██╗░███████║░╚██╗████╗██╔╝█████╗░░██████╦╝
	██║╚██╔╝██║██╔══██║██║╚████║██║░░╚██╗██╔══██║░░████╔═████║░██╔══╝░░██╔══██╗
	██║░╚═╝░██║██║░░██║██║░╚███║╚██████╔╝██║░░██║░░╚██╔╝░╚██╔╝░███████╗██████╦╝
	╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝░╚═════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚══════╝╚═════╝░
	Version: %s`, versionString)
	fmt.Println()
}
