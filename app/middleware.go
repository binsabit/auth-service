package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func (app *Application) JWT(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func (app *Application) Authorize(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))

	ok := app.Auth.IsAuthorized(context.Background(), user_id, user.Raw)

	if ok {
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "session has expired"})
}
