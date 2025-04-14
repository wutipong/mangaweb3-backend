package database

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	dialect_sql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/migrate"
	"github.com/wutipong/mangaweb3-backend/errors"
	_ "modernc.org/sqlite"
)

var pool *pgxpool.Pool
var sqlite *sql.DB

var db *dialect_sql.Driver
var databaseType string

func Open(ctx context.Context, dbType string, connStr string) error {
	databaseType = dbType

	switch dbType {
	case dialect.Postgres:
		return openPostgres(ctx, connStr)

	case dialect.SQLite:
		return openSqlite(ctx, connStr)
	}

	return errors.ErrNotImplemented

}

func openPostgres(ctx context.Context, connStr string) error {
	if p, e := pgxpool.New(ctx, connStr); e == nil {
		pool = p
		db = dialect_sql.OpenDB(dialect.Postgres, stdlib.OpenDBFromPool(pool))
	} else {
		return e
	}

	return nil
}

func openSqlite(ctx context.Context, connStr string) error {
	if d, e := sql.Open("sqlite", connStr); e == nil {
		sqlite = d
		db = dialect_sql.OpenDB(dialect.SQLite, d)
		return e
	} else {
		return e
	}
}

func Close() {
	if pool != nil {
		pool.Close()
	}

	if sqlite != nil {
		sqlite.Close()
	}
}

func CreateEntClient() *ent.Client {
	options := []ent.Option{
		ent.Driver(db),
	}

	config := configuration.Get()
	if config.DebugMode {
		options = append(options,
			ent.Debug(),
			ent.Log(func(params ...any) {
				if databaseType == "postgres" {
					stat := pool.Stat()

					log.Debug().
						Any("params", params).
						Int32("Acquired Conns", stat.AcquiredConns()).
						Int32("Idle Conns", stat.IdleConns()).
						Int32("Constructed Conns", stat.ConstructingConns()).
						Msg("Ent Debug")
				}
			}),
		)
	}

	client := ent.NewClient(options...)

	return client
}

func CreateSchema(ctx context.Context) error {
	client := CreateEntClient()
	defer client.Close()

	return client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true))
}
