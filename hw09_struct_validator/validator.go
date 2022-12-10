package hw09structvalidator

import (
	"errors"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

const tagValidateName = "validate"

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	structToValidate := reflect.TypeOf(v)
	if structToValidate.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	return nil
}
