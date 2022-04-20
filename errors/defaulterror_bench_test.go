package errors

import (
	"reflect"
	"testing"
)

func BenchmarkCreateDefaultError(b *testing.B) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *DefaultError
	}{
		{
			name: "with-message",
			args: args{message: "just some test"},
			want: &DefaultError{Message: "just some test"},
		},
		{
			name: "without-message",
			args: args{message: ""},
			want: &DefaultError{Message: ""},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := CreateDefaultError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				b.Errorf("CreateDefaultError() = %v, want %v", got, tt.want)
			}
		})
	}
}
