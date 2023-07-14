package app

import (
	"log"

	"github.com/binsabit/auth-service/config"
	"github.com/binsabit/auth-service/db"
	"github.com/gofiber/fiber/v2"
)

type Application struct {
	User   db.UserStorage
	Opt    db.OtpStorage
	Auth   db.TokenStorage
	Config config.Config
	Router *fiber.App
}

func NewApplication(user db.UserStorage, opt db.OtpStorage, auth db.TokenStorage, config config.Config, router *fiber.App) *Application {
	return &Application{
		User:   user,
		Opt:    opt,
		Router: router,
		Config: config,
		Auth:   auth,
	}
}

func (app Application) RunApp() {
	app.SetupRoutes()
	log.Fatal(app.Router.Listen(app.Config.HTTPServer.Port))
}
