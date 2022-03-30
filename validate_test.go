package validate

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func Test_createErrorArray(t *testing.T) {
	v := validator.New()

	user := User{}
	err := v.Struct(user)

	errorArray := createErrorArray(err)

	assert.Equal(t, 2, len(errorArray))
}

func TestRun(t *testing.T) {
	errors := Run(nil)

	assert.Nil(t, errors)
}

func TestRun_no_errors(t *testing.T) {
	user := User{
		FullName: "João da Silva",
		Contact: []Contact{
			{
				Type:  "personal",
				Email: "test@gmail.com",
			},
		},
	}

	errors := Run(user)

	assert.Nil(t, errors)
}

func TestRun_no_email_error(t *testing.T) {
	user := User{
		FullName: "João da Silva",
		Contact: []Contact{
			{
				Type:  "personal",
				Email: "not-email",
			},
		},
	}

	errors := Run(user)

	assert.Equal(t, 1, len(errors))
}

func TestRun_no_2_errors(t *testing.T) {
	user := User{
		FullName: "João da Silva",
		Contact: []Contact{
			{
				Type: "personal",
			},
		},
	}

	errors := Run(user)

	assert.Equal(t, 2, len(errors))
}
