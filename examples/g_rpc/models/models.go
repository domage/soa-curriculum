package models

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var PG *bun.DB

type dbLogger struct{}

func getKey(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

func (d dbLogger) BeforeQuery(c context.Context, q *bun.QueryEvent) context.Context {
	query := q.Query //.FormattedQuery()
	fmt.Println(string(query))
	fmt.Println("")

	return c
}

func (d dbLogger) AfterQuery(c context.Context, q *bun.QueryEvent) {
}

func Init() {
	host := getKey("DBHOST", "localhost")
	dsn := "postgres://brewone:brewone@" + host + ":5432/brewone?sslmode=disable"

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	PG = bun.NewDB(db, pgdialect.New())
	PG.AddQueryHook(dbLogger{})
}
