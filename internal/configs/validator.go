package configs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type (
	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationErrors struct {
		Errors []ValidationError
	}

	StructValidator struct {
		validate *validator.Validate
	}
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

func (ve *ValidationErrors) Error() string {
	return fmt.Sprintf("validation failed: %v errors", len(ve.Errors))
}

func NewValidator() fiber.StructValidator {
	return &StructValidator{validate: validator.New()}
}

func (v *StructValidator) Validate(data interface{}) error {
	err := v.validate.Struct(data)
	if err == nil {
		return nil
	}

	var validationErrors []ValidationError
	for _, fieldErr := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, ValidationError{
			Field:   fieldErr.Field(),
			Message: v.errorMessage(fieldErr),
		})
	}

	return &ValidationErrors{Errors: validationErrors}
}

func (v *StructValidator) errorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required."
	case "email":
		return "Invalid email address."
	default:
		return "Invalid value."
	}
}
