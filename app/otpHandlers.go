package app

import (
	"context"
	"errors"

	"github.com/binsabit/auth-service/db/postgres"
	"github.com/binsabit/auth-service/util/json"
	"github.com/binsabit/auth-service/util/validator"
	"github.com/gofiber/fiber/v2"
)

// OTP = one-time-password
func (app Application) GetOTP(ctx *fiber.Ctx) error {
	var reqOtp getOtp

	if err := json.ReadJSON(ctx, &reqOtp); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := validator.ValidateStruct(&reqOtp); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := validator.ValidatePhoneNumber(reqOtp.Phone); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	code, err := app.Opt.CreateOtp(ctx.Context(), reqOtp.Phone)
	if err != nil {
		if errors.Is(err, postgres.CodeAlreadyExists) {
			return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
		}
		return json.ErrorJSON(ctx, err, fiber.StatusInternalServerError)
	}
	return json.WriteJSON(ctx, fiber.StatusOK, resOpt{Code: code})
}

func (app Application) VerifyOTP(ctx *fiber.Ctx) error {

	var verReq verifyOtp
	if err := json.ReadJSON(ctx, &verReq); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	if err := validator.ValidateStruct(&verReq); err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	err := validator.ValidatePhoneNumber(verReq.Phone)
	if err != nil {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	verified, err := app.Opt.VerifyOtp(context.TODO(), verReq.Phone, verReq.Code)

	if err != nil || !verified {
		return json.ErrorJSON(ctx, err, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)

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
