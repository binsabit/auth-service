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

func (psql PGXToken) AddToAuthTable(ctx context.Context, user_id int, token string) error {
	query := `INSERT INTO auth_table (user_id, token) VALUES ($1,$2)`
	_, err := psql.Conn.Exec(ctx, query, user_id, GenereteHash(token))

	return err
}

func (psql PGXToken) DeleteFromAuthTable(ctx context.Context, user_id int) error {
	query := `DELETE FROM auth_table WHERE user_id = $1 AND auth_id = $2`

	_, err := psql.Conn.Exec(ctx, query, user_id)

	return err
}

func (psql PGXToken) DeleteIfExpired(ctx context.Context) error {
	query := `DELETE FROM auth_table WHERE expires_at<$1`

	_, err := psql.Conn.Exec(ctx, query, time.Now())

	return err
}

func (psql PGXToken) IsAuthorized(ctx context.Context, user_id int, testToken string) bool {
	query := `SELECT token FROM auth_table WHERE user_id = $1`

	var token []byte

	row := psql.Conn.QueryRow(ctx, query, user_id)

	err := row.Scan(&token)

	if err == pgx.ErrNoRows {
		return false
	}

	if string(token) == string(GenereteHash(testToken)) {
		return true
	}

	return false
}
