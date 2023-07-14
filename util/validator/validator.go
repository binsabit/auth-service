package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var (
	phoneRegexpKZ = `^\d{3}\d{3}\d{2}\d{2}$`
)

type ValidationError struct {
	Message string
	Field   string
	Tag     string
}

func (ve *ValidationError) Error() string {
	return ve.Message
}

func ValidatePhoneNumber(number string) error {
	if len(number) != 10 {
		return &ValidationError{
			Message: "Phone number should be equal to 10 digits",
			Field:   "phone",
		}
	}

	done, err := regexp.MatchString(phoneRegexpKZ, number)
	if err != nil {
		return err
	}
	if !done {
		return &ValidationError{
			Message: "Incorrect phone format for KZ",
			Field:   "phone",
		}
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return &ValidationError{
			Message: "Password should be of 8 characters long",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one lower case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one upper case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one digit",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one special character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	return nil

}

func ValidateStruct(data interface{}) error {

	return validate.Struct(data)

}
