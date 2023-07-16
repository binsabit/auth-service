package postgres

import (
	"context"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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

func RunMigrations(migrationFile string, dsn string) {
	m, err := migrate.New(migrationFile, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
