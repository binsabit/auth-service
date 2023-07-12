package postgres

import (
	"context"
	"crypto/sha256"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrorNotFound           = errors.New("user not found")
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

type PGXUser struct {
	Conn *pgxpool.Pool
}

func NewPGXuser(conn *pgxpool.Pool) PGXUser {
	return PGXUser{
		Conn: conn,
	}
}
func (psql PGXUser) CreateUser(ctx context.Context, args ...string) error {
	query := `INSERT INTO users (phone, password, firstname, lastname) 
			VALUES ($1,$2,$3,$4)
	`
	_, err := psql.Conn.Exec(ctx, query, args[0], genereteHash(args[1]), args[2], args[3])
	if err != nil {
		return err
	}

	return nil
}

func (psql PGXUser) CheckCredentials(ctx context.Context, phone, password string) error {
	query := `SELECT password FROM users WHERE phone = $1`

	row := psql.Conn.QueryRow(ctx, query, phone)
	var p string
	err := row.Scan(&p)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrorNotFound
		}
		return err
	}

	if string(genereteHash(password)) != string(p) {
		return ErrorInvalidCredentials
	}

	return nil

}
func (psql PGXUser) LogoutUser(ctx context.Context, id int) error {
	return nil

}

func genereteHash(text string) []byte {
	h := sha256.New()
	h.Write([]byte(text))
	res := h.Sum(nil)
	return res
}
