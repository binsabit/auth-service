package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (app *Application) GetExpirationTimeFromJWT(ctx *fiber.Ctx, key string) time.Duration {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	res := claims[key].(time.Duration)
	return res
}

func (app *Application) GenereteJWT() {

}
