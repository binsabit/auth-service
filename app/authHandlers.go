package app

import (
	"log"
	"time"

	"github.com/binsabit/auth-service/util/json"
	"github.com/binsabit/auth-service/util/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (app *Application) Test(ctx *fiber.Ctx) error {
	ctx.SendStatus(fiber.StatusOK)
	return nil
}

func (app *Application) Signup(ctx *fiber.Ctx) error {

	var req signupRequest

	if err := validator.ValidateBody(ctx, &req); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := validator.ValidatePhoneNumber(req.Phone); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := validator.ValidatePassword(req.Password); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := app.User.CreateUser(ctx.Context(), req.Phone, req.Password, req.Firstname, req.Lastname); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	ctx.SendStatus(fiber.StatusCreated)

	return nil
}

func (app *Application) Login(ctx *fiber.Ctx) error {

	var user loginRequest

	err := validator.ValidateBody(ctx, &user)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := validator.ValidatePhoneNumber(user.Phone); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}
	userID, err := app.User.CheckCredentials(ctx.Context(), user.Phone, user.Password)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	authID := uuid.New().String()
	expiratesAt := time.Now().Add(app.Config.JWT.Expires)
	log.Println(authID, expiratesAt)

	err = app.Auth.AddToAuthTable(ctx.Context(), userID, authID, expiratesAt)

	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	claims := jtoken.MapClaims{
		"user_id":   userID,
		"auth_uuid": authID,
		"exp":       expiratesAt.Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(app.Config.JWT.Secret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token": t,
	})
}

func (app *Application) Logout(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))
	auth_id := claims["auth_uuid"].(string)

	err := app.Auth.DeleteFromAuthTable(ctx.Context(), user_id, auth_id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

type signupRequest struct {
	Phone     string `json:"phone,omitempty" validate:"required"`
	Password  string `json:"password,omitempty" validate:"required"`
	Firstname string `json:"firstname,omitempty" validate:"required"`
	Lastname  string `json:"lastname,omitempty" validate:"required"`
}

type loginRequest struct {
	Phone    string `json:"phone,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
