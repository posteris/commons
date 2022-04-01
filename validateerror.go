package commons

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func newValidateError(err validator.FieldError) ValidationError {
	translator := createTranslator()

	var validationError ValidationError

	validationError.Field = err.StructNamespace()
	validationError.Message = err.Translate(translator)

	return validationError
}

func createTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}
