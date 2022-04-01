package commons

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

func getErrorArray(user User) []ValidationError {
	validate = validator.New()
	err := validate.Struct(user)

	if err == nil {
		return []ValidationError{}
	}

	var validationError []ValidationError

	for _, valueError := range err.(validator.ValidationErrors) {
		validationError = append(validationError, newValidateError(valueError))
	}

	return validationError
}

var testCase = []struct {
	Name   string
	Errors []ValidationError
	Want   []ValidationError
}{
	{
		Name:   "all-in",
		Errors: getErrorArray(User{Name: "João", Surname: "da Silva", Email: "dasilva@gmail.com"}),
		Want:   []ValidationError{},
	},
	{
		Name:   "no-name",
		Errors: getErrorArray(User{Surname: "da Silva", Email: "dasilva@gmail.com"}),
		Want:   []ValidationError{{Field: "User.Name", Message: "Name is a required field"}},
	},
	{
		Name:   "no-surname",
		Errors: getErrorArray(User{Name: "João", Email: "dasilva@gmail.com"}),
		Want:   []ValidationError{{Field: "User.Surname", Message: "Surname is a required field"}},
	},
	{
		Name:   "no-email",
		Errors: getErrorArray(User{Name: "João", Surname: "da Silva"}),
		Want:   []ValidationError{},
	},
	{
		Name:   "invalid-email",
		Errors: getErrorArray(User{Name: "João", Surname: "da Silva", Email: "not-email"}),
		Want:   []ValidationError{{Field: "User.Email", Message: "Email must be a valid email address"}},
	},
	{
		Name:   "invalid-email-&-short-name",
		Errors: getErrorArray(User{Name: "Jo", Surname: "da Silva", Email: "not-email"}),
		Want:   []ValidationError{{Field: "User.Name", Message: "Name must be at least 3 characters in length"}, {Field: "User.Email", Message: "Email must be a valid email address"}},
	},
}

func TestNewValidateError_qt(t *testing.T) {

	for _, tt := range testCase {
		t.Run(tt.Name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.Errors, tt.Want) {
				t.Errorf("Errors = %v, want %v", tt.Errors, tt.Want)
			}
		})
	}
}
