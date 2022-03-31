package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// CustomError for posteris platform
type CustomError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	// Value   string `json:"value,omitempty"`
}

//createCustomError function to creates a new CustomError
func createCustomError(err validator.FieldError) *CustomError {
	if err == nil {
		return nil
	}

	var customError CustomError

	customError.Field = err.StructNamespace()

	msg := err.Translate(translate)
	strFind := "Error:"
	index := strings.Index(msg, strFind)

	customError.Message = msg[(index + len(strFind)):]

	return &customError
}
