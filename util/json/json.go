package json

import "github.com/gofiber/fiber/v2"

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func WriteJSON(ctx *fiber.Ctx, status int, data any) error {
	return ctx.Status(status).JSON(jsonResponse{Status: status, Message: "success", Data: data})
}

func ErrorJSON(ctx *fiber.Ctx, err error, status ...int) error {
	statusCode := fiber.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	return ctx.Status(statusCode).JSON(jsonResponse{Status: statusCode, Message: err.Error()})
}
