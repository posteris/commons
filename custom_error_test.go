package validate

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type Contact struct {
	Type  string `validate:"required,oneof=commercial personal billing sales"`
	Email string `validate:"required_without=Phone,omitempty,email"`
	Phone string `validate:"required_without=Email,omitempty,numeric"`
}

type User struct {
	ID       string
	FullName string    `validate:"required,min=6,max=256" trans:"Full Name"`
	Contact  []Contact `validate:"gt=0,dive"`
	IsActive bool
}

func Test_createCustomError_empty_struct(t *testing.T) {
	v := validator.New()

	user := User{}
	err := v.Struct(user)

	count := 0
	for _, errElement := range err.(validator.ValidationErrors) {
		customError := createCustomError(errElement)

		if customError != nil {
			count = count + 1
		}
	}

	assert.Equal(t, 2, count)
}

func Test_createCustomError(t *testing.T) {
	v := validator.New()

	user := User{
		FullName: "João da Silva",
		Contact: []Contact{
			{
				Type: "personal",
			},
		},
	}
	err := v.Struct(user)

	count := 0
	for _, errElement := range err.(validator.ValidationErrors) {
		customError := createCustomError(errElement)

		if customError != nil {
			count = count + 1
		}
	}

	assert.Equal(t, 2, count)
}

func Test_createCustomError_complete_error(t *testing.T) {
	v := validator.New()

	user := User{
		FullName: "João da Silva",
		Contact: []Contact{
			{
				Type:  "personal",
				Email: "not-email",
			},
		},
	}
	err := v.Struct(user)

	for _, errElement := range err.(validator.ValidationErrors) {
		customError := createCustomError(errElement)

		expected := &CustomError{
			Field:   "User.Contact[0].Email",
			Message: "Field validation for 'Email' failed on the 'email' tag",
		}

		if customError != nil {
			assert.Equal(t, expected, customError)
		}
	}
}

func Test_createCustomError_nil(t *testing.T) {
	err := createCustomError(nil)

	assert.Nil(t, err)
}
