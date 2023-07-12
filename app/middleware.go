package app

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func (app *Application) Authenticate(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
