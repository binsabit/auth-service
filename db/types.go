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
	AddToBlackList(context.Context, string) error
	DeleteFromBlackList(context.Context, string) error
	IsInBlackList(context.Context, string) bool
}

type OtpStorage interface {
	CreateOtp(context.Context, string) (string, error)
	VerifyOtp(context.Context, string, string) (bool, error)
}

type UserStorage interface {
	CreateUser(context.Context, ...string) error
	CheckCredentials(context.Context, string, string) error
	LogoutUser(context.Context, int) error
}
