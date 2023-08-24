package main

import (
	"context"
	"flag"

	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/wutipong/mangaweb3-backend/docs"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/handler/browse"
	handlertag "github.com/wutipong/mangaweb3-backend/handler/tag"
	"github.com/wutipong/mangaweb3-backend/handler/view"
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
	pathBrowse         = "/browse"
	pathView           = "/view"
	pathGetImage       = "/get_image"
	pathUpdateCover    = "/update_cover"
	pathThumbnail      = "/get_thumbnail"
	pathFavorite       = "/favorite"
	pathDownload       = "/download"
	pathRescanLibrary  = "/rescan_library"
	pathSetTagFavorite = "/set_tag_favorite"
	pathTagList        = "/tag_list"
	pathTagThumb       = "/tag_thumb"
)

//go:generate swag init

// @title           Mangaweb3 API
// @version         3.0
// @description     API Server for Mangaweb

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Info().Msg("Use .env file.")
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

	log.Info().
		Str("version", versionString).
		Str("data_path", *dataPath).
		Str("address", *address).
		Msg("Server started.")

	router := httprouter.New()
	conn, err := pgxpool.New(context.Background(), *connectionStr)
	if err != nil {
		log.Error().AnErr("error", err).Msg("Connect to Postgres fails")

		return
	}

	defer conn.Close()

	tag.Init(conn)
	meta.Init(conn)

	scheduler.Init(scheduler.Options{})

	RegisterHandler(router)
	scheduler.Start()

	log.Info().Msg("Server starts.")
	if err = http.ListenAndServe(*address, router); err != nil {
		log.Error().AnErr("error", err).Msg("Starting server fails")
		return
	}

	log.Info().Msg("shutting down the server")
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
		PathTagFavorite:   pathSetTagFavorite,
		PathTagList:       pathTagList,
		PathTagThumbnail:  pathTagThumb,
	})
	// Routes
	router.POST(pathBrowse, browse.Handler)
	router.POST(pathView, view.Handler)
	router.GET(pathGetImage, handler.GetImage)
	router.POST(pathUpdateCover, view.UpdateCover)
	router.GET(pathThumbnail, browse.GetThumbnailHandler)
	router.POST(pathFavorite, view.SetFavoriteHandler)
	router.GET(pathDownload, view.Download)
	router.GET(pathRescanLibrary, handler.RescanLibraryHandler)
	router.POST(pathSetTagFavorite, handlertag.SetFavoriteHandler)
	router.GET(pathTagList, handlertag.TagListHandler)
	router.GET(pathTagThumb, handlertag.ThumbnailHandler)

	router.GET("/doc/:any", swaggerHandler)
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
