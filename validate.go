package validate

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate = validator.New()

func getTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}

var translate = getTranslator()

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
