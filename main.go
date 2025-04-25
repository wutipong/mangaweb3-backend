package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/wutipong/mangaweb3-backend/configuration"
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
	ctx := context.Background()

	useEnvFile := false
	if err := godotenv.Overload(); err == nil {
		useEnvFile = true
	}

	address := ":8972"
	if value, valid := os.LookupEnv("MANGAWEB_ADDRESS"); valid {
		address = value
	}

	dataPath := "./data"
	if value, valid := os.LookupEnv("MANGAWEB_DATA_PATH"); valid {
		dataPath = value
	}

	cachePath := "./cache"
	if value, valid := os.LookupEnv("MANGAWEB_CACHE_PATH"); valid {
		cachePath = value
	}

	connectionStr := "file:db.sqlite3?_pragma=foreign_keys(1)"
	if value, valid := os.LookupEnv("MANGAWEB_DB"); valid {
		connectionStr = value
	}

	dbType := dialect.SQLite
	if value, valid := os.LookupEnv("MANGAWEB_DB_TYPE"); valid {
		dbType = value
	}

	debugMode := false
	environment := "development"
	if value, valid := os.LookupEnv("MANGAWEB_ENVIRONMENT"); valid {
		environment = value
	}

	if strings.ToLower(strings.TrimSpace(environment)) == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
			Level(zerolog.DebugLevel)

		log.Info().Msg("development environment")
		debugMode = true
	}

	if useEnvFile {
		log.Info().Msg("Use .env file.")
	}

	log.Info().
		Bool("debugMode", debugMode).
		Str("version", versionString).
		Str("dataPath", dataPath).
		Str("cachePath", cachePath).
		Str("dbType", dbType).
		Str("dbConnection", connectionStr).
		Str("address", address).
		Msg("Server started.")

	configuration.Init(configuration.Config{
		DebugMode:     debugMode,
		VersionString: versionString,
		DataPath:      dataPath,
		CachePath:     cachePath,
	})

	if err := database.Open(ctx, dbType, connectionStr); err != nil {
		log.Error().AnErr("error", err).Msg("Connect to Database fails")
		return
	} else {
		defer database.Close()
	}

	if err := database.CreateSchema(ctx); err != nil {
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
