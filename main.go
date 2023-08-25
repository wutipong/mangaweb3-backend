package main

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"github.com/rs/cors"
	"github.com/rs/zerolog"
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

var versionString string = "development"

//go:generate swag init

// @title           Mangaweb3 API
// @version         3.0
// @description     API Server for Mangaweb

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Error().
			AnErr("error", err).
			Msg("Use .env file.")

		return
	}

	address := ":8972"
	if v, b := os.LookupEnv("MANGAWEB_ADDRESS"); b {
		address = v
	}

	dataPath := "./data"
	if v, b := os.LookupEnv("MANGAWEB_DATA_PATH"); b {
		dataPath = v
	}
	connectionStr := "postgres://postgres:password@localhost:5432/manga"
	if v, b := os.LookupEnv("MANGAWEB_DB"); b {
		connectionStr = v
	}

	if _, b := os.LookupEnv("MANGAWEB_DEVELOPMENT"); b {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
			Level(zerolog.DebugLevel)
	}

	meta.BaseDirectory = dataPath

	log.Info().
		Str("version", versionString).
		Str("data_path", dataPath).
		Str("address", address).
		Msg("Server started.")

	router := httprouter.New()
	conn, err := pgxpool.New(context.Background(), connectionStr)
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

	handler := cors.Default().Handler(router)
	if err = http.ListenAndServe(address, handler); err != nil {
		log.Error().AnErr("error", err).Msg("Starting server fails")
		return
	}

	log.Info().Msg("shutting down the server")
	scheduler.Stop()
}

func RegisterHandler(router *httprouter.Router) {
	handler.Init(handler.Options{
		VersionString:     versionString,
		PathBrowse:        browse.PathBrowse,
		PathView:          view.PathView,
		PathGetImage:      view.PathGetImage,
		PathUpdateCover:   view.PathUpdateCover,
		PathThumbnail:     browse.PathThumbnail,
		PathFavorite:      view.PathFavorite,
		PathDownload:      view.PathDownload,
		PathRescanLibrary: browse.PathRescanLibrary,
		PathTagFavorite:   handlertag.PathSetFavorite,
		PathTagList:       handlertag.PathList,
		PathTagThumbnail:  handlertag.PathThumbnail,
	})
	// Routes
	router.POST(browse.PathBrowse, browse.Handler)
	router.POST(view.PathView, view.Handler)
	router.GET(view.PathGetImage, view.GetImage)
	router.POST(view.PathUpdateCover, view.UpdateCover)
	router.GET(browse.PathThumbnail, browse.GetThumbnailHandler)
	router.POST(view.PathFavorite, view.SetFavoriteHandler)
	router.GET(view.PathDownload, view.Download)
	router.GET(browse.PathRescanLibrary, browse.RescanLibraryHandler)
	router.POST(handlertag.PathSetFavorite, handlertag.SetFavoriteHandler)
	router.POST(handlertag.PathList, handlertag.ListHandler)
	router.GET(handlertag.PathThumbnail, handlertag.ThumbnailHandler)

	router.GET("/doc/:any", swaggerHandler)
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
