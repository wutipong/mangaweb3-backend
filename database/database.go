package database

import (
	"context"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
)

var pool *pgxpool.Pool
var debugMode = false

func Open(ctx context.Context, connStr string) error {
	if p, e := pgxpool.New(ctx, connStr); e == nil {
		pool = p
	} else {
		return e
	}

	return nil
}

func Close() {
	pool.Close()
}

func CreateEntClient() *ent.Client {
	drv := sql.OpenDB(dialect.Postgres, stdlib.OpenDBFromPool(pool))
	options := []ent.Option{
		ent.Driver(drv),
		ent.Log(func(params ...any) {
			log.Debug().Any("params", params).Msg("Ent Debug")
		}),
	}
	if debugMode {
		options = append(options, ent.Debug())
	}

	client := ent.NewClient(options...)

	return client
}

func CreateSchema() error {
	client := CreateEntClient()
	defer client.Close()

	return client.Schema.Create(context.Background())
}
