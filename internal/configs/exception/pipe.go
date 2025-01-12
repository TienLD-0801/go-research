package exception_configs

import (
	"go-backend/internal/constants"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

const (
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrRequestTimeout      = "Request Timeout"
	ErrInvalidEmail        = "Invalid email"
	ErrInvalidPassword     = "Invalid password"
	ErrInvalidField        = "Invalid field"
	ErrInternalServerError = "Internal Server Error"
)

type (
	RestError struct {
		ErrStatus  int         `json:"status,omitempty"`
		ErrError   string      `json:"error,omitempty"`
		ErrMessage interface{} `json:"message,omitempty"`
		Timestamp  time.Time   `json:"timestamp,omitempty"`
	}

	RestErr interface {
		Status() int
		Error() string
		Causes() interface{}
		ErrBody() RestError
	}

	CustomValidator struct {
		validate *validator.Validate
	}
)

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validate: validator.New(),
	}
}

func (v CustomValidator) Validate(data interface{}) []RestError {
	var validationErrors []RestError

	errs := v.validate.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var ve RestError

			ve.ErrError = err.Field()
			ve.ErrMessage = err.Tag()
			ve.ErrStatus = 400

			validationErrors = append(validationErrors, ve)
		}
	}

	return validationErrors
}

func SetCustomValidatorContext(c fiber.Ctx) error {
	cv := NewCustomValidator()
	c.Locals(constants.CustomValidator, cv)
	return c.Next()
}
