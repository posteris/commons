package commons

import (
	"github.com/go-playground/validator/v10"
)

//createErrorArray function to create an CustomError array based on an
//error received by parameter
func createErrorArray(err error) []ValidationError {
	var errorArray []ValidationError

	for _, errElement := range err.(validator.ValidationErrors) {
		customError := newValidateError(errElement)
		errorArray = append(errorArray, customError)
	}

	return errorArray
}

//Run function to perform the struct validate. This function returns an array
//or errors or nil.
func ValidateModel(model interface{}) []ValidationError {
	if model == nil {
		return []ValidationError{}
	}

	err := validate.Struct(model)

	if err != nil {
		array := createErrorArray(err)

		if len(array) > 0 {
			return array
		}
	}

	return []ValidationError{}
}
