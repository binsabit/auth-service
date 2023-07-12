package postgres

import (
	"context"
	"log"

	// "github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection(ctx context.Context, dsn string) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("could not connect to database:%s \n error:%v", dsn, err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("could not ping database:%s \n error:%v", dsn, err)
	}
	return conn
}
