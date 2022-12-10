package hw09structvalidator

import (
	"errors"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

const (
	tagValidateName       = "validate"
	rulesSplitter         = "|"
	ruleNameValueSplitter = ":"
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	structToValidate := reflect.TypeOf(v)
	if structToValidate.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	for i := 0; i < structToValidate.NumField(); i++ {
		field := structToValidate.Field(i)
		tag := field.Tag.Get(tagValidateName)

		// skip fields without "validate" tag
		if len(tag) == 0 {
			continue
		}

		// Here we split rules by "|" seprator and next split each rule
		// by key and value with ":" separator
		// ex. `validate:"min:18|max:50"`
		for _, value := range strings.Split(tag, rulesSplitter) {
			rule := strings.SplitN(value, ruleNameValueSplitter, 2)
			if len(rule) != 2 {
				return errors.New("wrong type of rule")
			}
		}
	}
	return nil
}
