package commons

type DefaultError struct {
	Message string `json:"message"`
}

//CreateDefaultError function to generate new error message
func CreateDefaultError(message string) *DefaultError {
	return &DefaultError{
		Message: message,
	}
}
