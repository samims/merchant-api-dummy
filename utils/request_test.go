package utils

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestCheckAndSetString(t *testing.T) {
	// Arrange
	existingField := fake.Word()
	newField := fake.Word()

	type args struct {
		existingField string
		newField      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Test cases.

		{
			name: "existing field is empty",
			args: args{
				existingField: "",
				newField:      newField,
			},
			want: newField,
		},
		{
			name: "new field is empty",
			args: args{
				existingField: existingField,
				newField:      "",
			},
			want: existingField,
		},
		{
			name: "both fields are not empty",
			args: args{
				existingField: existingField,
				newField:      newField,
			},
			want: newField,
		},
		{
			name: "both fields are empty",
			args: args{
				existingField: "",
				newField:      "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckAndSetString(tt.args.existingField, tt.args.newField); got != tt.want {
				t.Errorf("CheckAndSetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
