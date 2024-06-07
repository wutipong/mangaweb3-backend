package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/wutipong/mangaweb3-backend/docs"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/handler/browse"
	handlertag "github.com/wutipong/mangaweb3-backend/handler/tag"
	"github.com/wutipong/mangaweb3-backend/handler/view"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/scheduler"
	"goji.io"
	"goji.io/pat"
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

	debugMode := false
	if v, b := os.LookupEnv("MANGAWEB_ENVIRONMENT"); b {
		if strings.ToLower(strings.TrimSpace(v)) == "development" {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
				Level(zerolog.DebugLevel)

			log.Info().Msg("development environment")
			debugMode = true
		}
	}

	meta.BaseDirectory = dataPath

	log.Info().
		Str("version", versionString).
		Str("data_path", dataPath).
		Str("address", address).
		Msg("Server started.")

	var client *ent.Client = nil
	if db, err := sql.Open("pgx", connectionStr); err != nil {
		log.Error().AnErr("error", err).Msg("Connect to Postgres fails")
		return
	} else {
		drv := entsql.OpenDB(dialect.Postgres, db)
		defer db.Close()

		options := []ent.Option{
			ent.Driver(drv),
			ent.Log(func(params ...any) {
				log.Debug().Any("params", params).Msg("Ent Debug")
			}),
		}
		if debugMode {
			options = append(options, ent.Debug())
		}

		client = ent.NewClient(options...)
		defer client.Close()
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Error().AnErr("error", err).Msg("failed creating schema resources.")
		return
	}

	scheduler.Init(scheduler.Options{
		EntClient: client,
	})
	scheduler.Start()

	log.Info().Msg("Server starts.")

	mux := goji.NewMux()
	RegisterHandler(mux, client)

	c := cors.AllowAll()
	mux.Use(c.Handler)

	if err := http.ListenAndServe(address, mux); err != nil {
		log.Error().AnErr("error", err).Msg("Starting server fails")
		return
	}

	log.Info().Msg("shutting down the server")
	scheduler.Stop()
}

func RegisterHandler(mux *goji.Mux, client *ent.Client) {
	handler.Init(handler.Options{
		VersionString: versionString,
		EntClient:     client,
	})

	browse.Register(mux)
	handlertag.Register(mux)
	view.Register(mux)

	mux.Handle(pat.Get("/doc/:any"), httpSwagger.WrapHandler)
}
