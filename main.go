package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
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
	"github.com/wutipong/mangaweb3-backend/service"
	"github.com/wutipong/mangaweb3-backend/service/impl"
	"github.com/wutipong/mangaweb3-backend/tag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	meta.Init(client)
	tag.Init(client)

	scheduler.Init(scheduler.Options{})
	scheduler.Start()

	router := httprouter.New()
	RegisterHandler(router)

	log.Info().Msg("Server starts.")

	handler := cors.Default().Handler(router)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Error().AnErr("error", err).Msg("Listen gRPC error.")
	}

	grpcServer := grpc.NewServer()
	service.RegisterMetaServiceServer(
		grpcServer,
		&impl.MetaServer{
			EntClient: client,
		})

	reflection.Register(grpcServer)

	ch := make(chan error, 2)

	go func() {
		log.Info().Msg("listening REST API")
		ch <- http.ListenAndServe(address, handler)
	}()

	go func() {
		log.Info().Msg("listening gRPC")
		ch <- grpcServer.Serve(listener)
	}()

	err = <-ch

	log.Error().AnErr("error", err).Msg("Error")

	log.Info().Msg("shutting down the server")
	scheduler.Stop()
}

func RegisterHandler(router *httprouter.Router) {
	handler.Init(handler.Options{
		VersionString: versionString,
	})

	browse.Register(router)
	handlertag.Register(router)
	view.Register(router)

	router.GET("/doc/:any", swaggerHandler)
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
