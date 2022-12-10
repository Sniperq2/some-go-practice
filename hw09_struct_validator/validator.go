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

type Rule struct {
	Name  string
	Value string
}

type Rules []Rule

// Here we split rules by "|" seprator and next split each rule
// by key and value with ":" separator
// ex. `validate:"min:18|max:50"`
// rulesSep - separator between rules, ex. |
// ruleSep - separator between rule and value, ex. :.
func ParsingRules(tag string, rulesSep string, ruleSep string) (Rules, error) {
	rulesSoup := strings.Split(tag, rulesSep)

	rules := make(Rules, 0)
	for _, value := range rulesSoup {
		rule := strings.SplitN(value, ruleSep, 2)

		if len(rule) != 2 {
			return nil, errors.New("wrong type of rule")
		}

		rules = append(rules, Rule{
			Name:  rule[0],
			Value: rule[1],
		})
	}

	if len(rules) == 0 {
		return nil, errors.New("no rules found")
	}

	return rules, nil
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

		_, err := ParsingRules(tag, rulesSplitter, ruleNameValueSplitter)
		if err != nil {
			return err
		}
	}
	return nil
}
