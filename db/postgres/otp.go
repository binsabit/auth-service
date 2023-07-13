package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/binsabit/auth-service/util/otp"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	CodeExpired       = errors.New("code has been expired")
	CodeAlreadyExists = errors.New("one time password already has been sent")
	CodeIncorrect     = errors.New("user sent incorrect code")
)

type PGXOtp struct {
	Conn    *pgxpool.Pool
	Expires time.Duration
	Length  int
}

func NewPGXOtp(conn *pgxpool.Pool, expires time.Duration, length int) PGXOtp {
	return PGXOtp{
		Conn:    conn,
		Expires: expires,
		Length:  length,
	}
}

func (psql PGXOtp) CreateOtp(ctx context.Context, phone string) (string, error) {
	// query := psql.Conn.Prepare(`Insert `)
	exist, err := psql.Sent(ctx, phone)
	if err != nil || exist {
		return "", err
	}

	query := `INSERT INTO otps (phone,code,expires_at) VALUES ($1,$2,$3)`
	code := otp.GenereteOtpNum(psql.Length)

	_, err = psql.Conn.Exec(ctx, query, phone, code, time.Now().Add(psql.Expires))
	if err != nil {
		return "", err
	}

	return code, nil
}

func (psql PGXOtp) Exists() (bool, error) {

}

func (psql PGXOtp) Sent(ctx context.Context, phone string) (bool, error) {
	query := `SELECT id FROM otps WHERE phone=$1 AND expires_at>$2`

	row := psql.Conn.QueryRow(ctx, query, phone, time.Now())

	err := row.Scan(nil)

	switch {
	case err == pgx.ErrNoRows:
		return false, nil
	case err != nil:
		return false, err
	}

	return true, CodeAlreadyExists

}

func (psql PGXOtp) VerifyOtp(ctx context.Context, phone, code string) (bool, error) {

	query := `SELECT expires_at FROM otps
				WHERE
				phone = $1 
				AND
				code = $2
				ORDER BY expires_at DESC
	`

	row := psql.Conn.QueryRow(ctx, query, phone, code)
	var codeCreatedAt time.Time
	err := row.Scan(&codeCreatedAt)

	switch {
	case err == pgx.ErrNoRows:
		return false, CodeIncorrect
	case err != nil:
		return false, err
	case time.Since(codeCreatedAt) >= psql.Expires:
		return false, CodeIncorrect
	}

	return true, nil
}
