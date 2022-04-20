package errors

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func CreateValidationError(field string, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}
