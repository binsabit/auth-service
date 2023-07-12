package app

import (
	"context"
	"errors"
	"time"

	"github.com/binsabit/auth-service/db/postgres"
	"github.com/binsabit/auth-service/util/json"
	"github.com/binsabit/auth-service/util/validator"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

//OTP = one-time-password
func (app *Application) GetOTP(ctx *fiber.Ctx) error {
	var reqOtp getOtp

	err := validator.ValidateBody(ctx, &reqOtp)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	phone, err := validator.ValidatePhoneNumber(reqOtp.Phone)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	reqOtp.Phone = phone

	code, err := app.Opt.CreateOtp(context.Background(), reqOtp.Phone)
	if err != nil {
		if errors.Is(err, postgres.CodeAlreadyExists) {
			return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
		}
		return json.ErrorJSON(ctx, err, fiber.StatusInternalServerError)
	}

	return json.WriteJSON(ctx, fiber.StatusOK, resOpt{Code: code})
}

func (app *Application) VerifyOTP(ctx *fiber.Ctx) error {

	var verReq verifyOtp
	if err := validator.ValidateBody(ctx, &verReq); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	phone, err := validator.ValidatePhoneNumber(verReq.Phone)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}
	verReq.Phone = phone

	verified, err := app.Opt.VerifyOtp(context.TODO(), verReq.Phone, verReq.Code)

	if err != nil || !verified {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	ctx.SendStatus(fiber.StatusOK)

	return nil
}

func (app *Application) Signup(ctx *fiber.Ctx) error {

	var req signupRequest

	if err := validator.ValidateBody(ctx, &req); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	phone, err := validator.ValidatePhoneNumber(req.Phone)

	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}
	req.Phone = phone

	if err := app.User.CreateUser(context.Background(), req.Phone, req.Password, req.Firstname, req.Lastname); err != nil {
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

	phone, err := validator.ValidatePhoneNumber(user.Phone)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}
	user.Phone = phone

	if err := app.User.CheckCredentials(context.Background(), user.Phone, user.Password); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	claims := jtoken.MapClaims{
		"exp": time.Now().Add(app.Config.JWT.Expires).Unix(),
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

	app.GetExpirationTimeFromJWT(ctx, "exp")

	ctx.SendStatus(fiber.StatusOK)

	return nil
}

type getOtp struct {
	Phone string `json:"phone,omitempty" validate:"required"`
}

type resOpt struct {
	Code string `json:"code,omitempty" validate:"required"`
}

type verifyOtp struct {
	Phone string `json:"phone,omitempty" validate:"required"`
	Code  string `json:"code,omitempty" validate:"required"`
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
