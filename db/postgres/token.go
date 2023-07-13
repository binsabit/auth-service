package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXToken struct {
	Conn *pgxpool.Pool
}

func NewPGXToken(conn *pgxpool.Pool) PGXToken {
	return PGXToken{
		Conn: conn,
	}
}

func (psql PGXToken) AddToAuthTable(ctx context.Context, user_id int, auth_id string, expiresAt time.Time) error {
	query := `INSERT INTO auth_table (user_id, auth_id, expires_at) VALUES ($1,$2,$3)`
	_, err := psql.Conn.Exec(ctx, query, user_id, auth_id, expiresAt)

	return err
}

func (psql PGXToken) DeleteFromAuthTable(ctx context.Context, auth_id, user_id int) error {
	query := `DELETE FROM auth_table WHERE user_id = $1 && auth_id = $2`

	_, err := psql.Conn.Exec(ctx, query, auth_id, user_id)

	return err
}

func (psql PGXToken) DeleteIdExpired(ctx context.Context) error {
	query := `DELETE FROM auth_table WHERE expires_at<$1`

	_, err := psql.Conn.Exec(ctx, query, time.Now())

	return err
}

func (psql PGXToken) IsAuthorized(ctx context.Context, user_id, auth_id int) bool {
	query := `SELECT id FROM auth_table WHERE user_id = $1 auth_id = $2 AND expires_at<=$3`

	var id int

	row := psql.Conn.QueryRow(ctx, query, user_id, auth_id, time.Now())

	err := row.Scan(&id)

	if err == pgx.ErrNoRows {
		return false
	}

	if id > 0 {
		return true
	}

	return err == nil
}
