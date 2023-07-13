package db

import (
	"context"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Phone     string `json:"phone"`
	Password  []byte `json:"password"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}

type Otp struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Phone     string    `json:"phone"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type TokenStorage interface {
	AddToAuthTable(context.Context, int, string, time.Time) error
	DeleteFromAuthTable(context.Context, int, int) error
	DeleteIdExpired(context.Context) error
	IsAuthorized(context.Context, int, int) bool
}

type OtpStorage interface {
	CreateOtp(context.Context, string) (string, error)
	VerifyOtp(context.Context, string, string) (bool, error)
}

type UserStorage interface {
	CreateUser(context.Context, ...string) error
	CheckCredentials(context.Context, string, string) (int, error)
	LogoutUser(context.Context, int) error
}
