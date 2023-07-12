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

func (psql PGXOtp) AddToBlackList(ctx context.Context, token string, expiresAt time.Time) error {
	query := `INSERT INTO jwt_blacklist (token, expires_at) VALUES ($1,$2)`

	_, err := psql.Conn.Exec(ctx, query, token, expiresAt)

	return err
}
func (psql PGXOtp) DeleteFromBlackList(ctx context.Context, token string) error {
	query := `DELETE FROM jwt_blacklist WHERE token = $1`

	_, err := psql.Conn.Exec(ctx, query, token)

	return err

}
func (psql PGXOtp) IsInBlackList(ctx context.Context, token string) bool {
	query := `SELECT id FROM jwt_blacklist WHERE token = $1`

	var id int

	row := psql.Conn.QueryRow(ctx, query, token)

	err := row.Scan(&id)

	if err == pgx.ErrNoRows {
		return false
	}

	if id > 0 {
		return true
	}

	return err == nil
}
