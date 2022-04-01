package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/posteris/custom-validate/errors"
)

var validate = validator.New()

func createTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}

//createErrorArray function to create an CustomError array based on an
//error received by parameter
func createErrorArray(err error) []errors.ValidationError {
	var errorArray []errors.ValidationError

	for _, errElement := range err.(validator.ValidationErrors) {
		translator := createTranslator()

		customError := errors.CreateValidationError(
			errElement.StructNamespace(),
			errElement.Translate(translator),
		)

		// customError := newValidateError(errElement)
		errorArray = append(errorArray, *customError)
	}

	return errorArray
}

//Run function to perform the struct validate. This function returns an array
//or errors or nil.
func ValidateModel(model interface{}) []errors.ValidationError {
	if model == nil {
		return []errors.ValidationError{}
	}

	err := validate.Struct(model)

	if err != nil {
		array := createErrorArray(err)

		if len(array) > 0 {
			return array
		}
	}

	return []errors.ValidationError{}
}
