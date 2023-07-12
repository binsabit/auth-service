package validator

import (
	"fmt"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

var phoneRegexp = `^(?:\+7|8)?\(?\d{3}\)?\d{3}?\d{2}?\d{2}$`

func ValidatePhoneNumber(number string) (string, error) {
	if len(number) > 12 {
		return "", fmt.Errorf("invalid phone number too long: %d must be <=12", len(number)-12)
	}

	re, err := regexp.Compile(phoneRegexp)
	if err != nil {
		log.Fatalf("regexp:%v", err)
	}

	if re.MatchString(number) {
		return number[len(number)-10:], nil
	}

	return "", fmt.Errorf("not a valid phone number")
}

func ValidateBody(ctx *fiber.Ctx, data interface{}) error {
	if err := ctx.BodyParser(data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
