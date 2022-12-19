package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:6"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			User{
				ID:     "765849",
				Age:    18,
				Email:  "admin@admin.com",
				Role:   UserRole("admin"),
				Phones: []string{"79991232233"},
			},
			nil,
		},
		{
			App{
				Version: "10000",
			},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			if err != nil {
				if validationError, ok := err.(ValidationErrors); ok {
					expectedError := tt.expectedErr.(ValidationErrors)
					assert.True(t, len(expectedError) == len(validationError), "amount of errors are right")
				} else {
					assert.True(t, errors.Is(err, tt.expectedErr), "same error")
				}
			}
			assert.Equal(t, err, tt.expectedErr, "validated")
		})
	}
}
