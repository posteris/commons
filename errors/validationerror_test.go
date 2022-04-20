package errors

import (
	"reflect"
	"testing"
)

func TestCreateValidationError(t *testing.T) {
	type args struct {
		field   string
		message string
	}
	tests := []struct {
		name string
		args args
		want *ValidationError
	}{
		{
			name: "basic-error",
			args: args{field: "test", message: "just some test"},
			want: &ValidationError{Field: "test", Message: "just some test"},
		},
		{
			name: "no-message",
			args: args{field: "test"},
			want: &ValidationError{Field: "test", Message: ""},
		},
		{
			name: "no-field",
			args: args{message: "just some test"},
			want: &ValidationError{Field: "", Message: "just some test"},
		},
		{
			name: "no-args",
			args: args{},
			want: &ValidationError{Field: "", Message: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateValidationError(tt.args.field, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
