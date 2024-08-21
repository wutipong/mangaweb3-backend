package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/wutipong/mangaweb3-backend/config"
	"github.com/wutipong/mangaweb3-backend/database"
	_ "github.com/wutipong/mangaweb3-backend/docs"
	"github.com/wutipong/mangaweb3-backend/handler/browse"
	maintenanceHandler "github.com/wutipong/mangaweb3-backend/handler/maintenance"
	"github.com/wutipong/mangaweb3-backend/handler/tag"
	"github.com/wutipong/mangaweb3-backend/handler/view"
	"github.com/wutipong/mangaweb3-backend/maintenance"
)

var versionString string = "development"

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag@latest init

// @title           Mangaweb3 API
// @version         3.0
// @description     API Server for Mangaweb

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Overload(); err == nil {
		log.Info().Msg("Use .env file.")
	}

	address := ":8972"
	if value, valid := os.LookupEnv("MANGAWEB_ADDRESS"); valid {
		address = value
	}

	dataPath := "./data"
	if value, valid := os.LookupEnv("MANGAWEB_DATA_PATH"); valid {
		dataPath = value
	}
	connectionStr := "postgres://postgres:password@localhost:5432/manga"
	if value, valid := os.LookupEnv("MANGAWEB_DB"); valid {
		connectionStr = value
	}

	debugMode := false
	if value, valid := os.LookupEnv("MANGAWEB_ENVIRONMENT"); valid {
		if strings.ToLower(strings.TrimSpace(value)) == "development" {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
				Level(zerolog.DebugLevel)

			log.Info().Msg("development environment")
			debugMode = true
		}
	}

	log.Info().
		Str("version", versionString).
		Str("data_path", dataPath).
		Str("address", address).
		Msg("Server started.")

	config.Init(config.Config{
		DebugMode:     debugMode,
		VersionString: versionString,
		DataPath:      dataPath,
	})

	if err := database.Open(context.Background(), connectionStr); err != nil {
		log.Error().AnErr("error", err).Msg("Connect to Postgres fails")
		return
	} else {
		defer database.Close()
	}

	if err := database.CreateSchema(); err != nil {
		log.Error().AnErr("error", err).Msg("failed creating schema resources.")
		return
	}

	go maintenance.UpdateLibrary()

	router := httprouter.New()
	RegisterHandler(router)

	log.Info().Msg("Server starts.")

	handler := cors.AllowAll().Handler(router)
	if err := http.ListenAndServe(address, handler); err != nil {
		log.Error().AnErr("error", err).Msg("Starting server fails")
		return
	}

	log.Info().Msg("shutting down the server")
}

func RegisterHandler(router *httprouter.Router) {
	browse.Register(router)
	tag.Register(router)
	view.Register(router)
	maintenanceHandler.Register(router)

	router.GET("/doc/:any", swaggerHandler)
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
