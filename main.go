package main

import (
	"context"
	"fmt"
	"log"

	"github.com/binsabit/auth-service/app"
	"github.com/binsabit/auth-service/config"
	"github.com/binsabit/auth-service/db/postgres"
	"github.com/binsabit/auth-service/sms"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.MustLoadConfig()

	log.Println(config)

	smsc := sms.NewSmsSender(config.Smsc.Login, config.Smsc.Password)

	storageConfig := config.Storage
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		storageConfig.Driver, storageConfig.User, storageConfig.Password, storageConfig.Host, storageConfig.Port, storageConfig.DBname, storageConfig.SSLMode,
	)

	//init storage
	pool := postgres.CreateConnection(context.Background(), dsn)
	usersStore := postgres.NewPGXuser(pool)
	optStore := postgres.NewPGXOtp(pool, config.OTP.Expires, config.OTP.Length)
	authStore := postgres.NewPGXToken(pool)

	//init router
	router := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	App := app.NewApplication(usersStore, optStore, authStore, smsc, config, router)

	App.RunApp()
}
