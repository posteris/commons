package validate

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

//createErrorArray function to create an CustomError array based on an
//error received by parameter
func createErrorArray(err error) []CustomError {
	var errorArray []CustomError

	for _, errElement := range err.(validator.ValidationErrors) {
		customError := createCustomError(errElement)
		errorArray = append(errorArray, *customError)
	}

	return errorArray
}

//Run function to perform the struct validate. This function returns an array
//or errors or nil.
func Run(model interface{}) []CustomError {
	if model == nil {
		return nil
	}

	err := validate.Struct(model)

	if err != nil {
		array := createErrorArray(err)

		if len(array) > 0 {
			return array
		}
	}

	return nil
}
