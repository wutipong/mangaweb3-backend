package database

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/ent"
)

var pool *pgxpool.Pool

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
	}

	config := configuration.Get()
	if config.DebugMode {
		options = append(options,
			ent.Debug(),
			ent.Log(func(params ...any) {
				stat := pool.Stat()

				log.Debug().
					Any("params", params).
					Int32("Acquired Conns", stat.AcquiredConns()).
					Int32("Idle Conns", stat.IdleConns()).
					Int32("Constructed Conns", stat.ConstructingConns()).
					Msg("Ent Debug")
			}),
		)
	}

	client := ent.NewClient(options...)

	return client
}

func CreateSchema(ctx context.Context) error {
	client := CreateEntClient()
	defer client.Close()

	return client.Schema.Create(
		ctx,
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				fmt.Println("Tables", tables)
				return next.Create(ctx, tables...)
			})
		}),
	)
}
