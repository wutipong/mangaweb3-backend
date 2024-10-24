package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

	if useEnvFile {
		log.Info().Msg("Use .env file.")
	}

	endpoint := "localhost:9000"
	if value, valid := os.LookupEnv("MANGAWEB_MINIO_ENDPOINT"); valid {
		endpoint = value
	}
	accessKey := ""
	if value, valid := os.LookupEnv("MANGAWEB_MINIO_ACCESS_KEY_ID"); valid {
		accessKey = value
	}

	accessSecret := ""
	if value, valid := os.LookupEnv("MANGAWEB_MINIO_ACCESS_KEY_SECRET"); valid {
		accessSecret = value
	}

	secure := false
	if value, valid := os.LookupEnv("MANGAWEB_MINIO_SECURE"); valid {
		if b, err := strconv.ParseBool(value); err != nil {
			log.Warn().Msg("Unable to parse MANGAWEB_MINIO_SECURE value. Default to false.")
			secure = false
		} else {
			secure = b
		}

	}

	bucket := "manga"
	if value, valid := os.LookupEnv("MANGAWEB_MINIO_BUCKET"); valid {
		bucket = value
	}

	var minioClient *minio.Client
	if c, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: secure,
	}); err != nil {
		log.Fatal().Msg("Unable to create MinIO client.")
		return
	} else {
		minioClient = c
	}

	if found, err := minioClient.BucketExists(context.Background(), bucket); !found || err != nil {
		if !found {
			log.Fatal().Msg("MinIO bucket not found.")
			return
		}

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("MinIO error.")

			return
		}
	}

	log.Info().
		Bool("debugMode", debugMode).
		Str("version", versionString).
		Str("data_path", dataPath).
		Str("address", address).
		Str("minio_endpoint", endpoint).
		Str("minio_access_key", accessKey).
		Bool("minio_secure", secure).
		Msg("Server started.")

	configuration.Init(configuration.Config{
		DebugMode:           debugMode,
		VersionString:       versionString,
		DataPath:            dataPath,
		MinIoEndPoint:       endpoint,
		MinIoAccessKey:      accessKey,
		MinIoAcessKeySecret: accessSecret,
		MinIoSecure:         secure,
		MinIoBucket:         bucket,
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
