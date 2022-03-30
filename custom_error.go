package validate

import "github.com/go-playground/validator/v10"

// CustomError for posteris platform
type CustomError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

//createCustomError function to creates a new CustomError
func createCustomError(err validator.FieldError) *CustomError {
	if err == nil {
		return nil
	}

	var customError CustomError

	customError.Field = err.StructNamespace()
	customError.Tag = err.Tag()

	value := err.Param()
	if value != "" {
		customError.Value = value
	}

	return &customError
}
