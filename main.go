package main

import (
	"context"
	"fmt"
	"log"

	"github.com/binsabit/auth-service/app"
	"github.com/binsabit/auth-service/config"
	"github.com/binsabit/auth-service/db/postgres"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.MustLoadConfig()

	log.Println(config)
	storageConfig := config.Storage
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		storageConfig.Driver, storageConfig.User, storageConfig.Password, storageConfig.Host, storageConfig.Port, storageConfig.DBname, storageConfig.SSLMode,
	)

	//init storage
	pool := postgres.CreateConnection(context.Background(), dsn)
	usersStore := postgres.NewPGXuser(pool)
	optStore := postgres.NewPGXOtp(pool, config.OTP.Expires, config.OTP.Length)

	//init router
	router := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	App := app.NewApplication(usersStore, optStore, config, router)

	App.RunApp()
}
