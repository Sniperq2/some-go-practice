package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrValueIsGreater         = errors.New("is greated than constraint")
	ErrUnsupportedRule        = errors.New("unsupported validation rule")
	ErrWrongRule              = errors.New("wrong rule")
	ErrBadLengthOf            = errors.New("bad length of")
	ErrNoInValueFound         = errors.New("no \"in\" value found")
	ErrCouldNotCompileRegExp  = errors.New("could not compile regexp value")
	ErrWrongRegularExpression = errors.New("wrong regular expression")
	ErrValueIsLesser          = errors.New("is lesser than constraint")
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
	return fmt.Sprintf("validation failed - number of errors is/are %d", len(v))
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
func parsingRules(tag string, rulesSep string, ruleSep string) (Rules, error) {
	rules := make(Rules, 0)
	for _, value := range strings.Split(tag, rulesSep) {
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

func inSlice(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func stringTypeConstraint(rules Rules, value reflect.Value) error {
	strValue := value.String()
	for _, rule := range rules {
		switch rule.Name {
		case "len":
			lengthValue, err := strconv.Atoi(rule.Value)
			if err != nil {
				return fmt.Errorf("%w %s - validation failed", ErrWrongRule, strValue)
			}

			if lengthValue != len(strValue) {
				return fmt.Errorf("%w %s - validation failed", ErrBadLengthOf, strValue)
			}
		case "in":
			inSplitted := strings.Split(rule.Value, ",")
			if !inSlice(inSplitted, strValue) {
				return fmt.Errorf("%w in %s - validation vailed", ErrNoInValueFound, strValue)
			}
		case "regexp":
			re, err := regexp.Compile(rule.Value)
			if err != nil {
				return fmt.Errorf("%w %s - validation vailed", ErrCouldNotCompileRegExp, strValue)
			}

			if !re.MatchString(strValue) {
				return fmt.Errorf("%w %s - validation vailed", ErrWrongRegularExpression, strValue)
			}
		default:
			return fmt.Errorf("%w %s - validation vailed", ErrUnsupportedRule, strValue)
		}
	}
	return nil
}

func validateMinMax(rules Rules, value reflect.Value) error {
	intValue := value.Int()
	for _, rule := range rules {
		switch rule.Name {
		case "min":
			min, err := strconv.Atoi(rule.Value)
			if err != nil {
				return fmt.Errorf("%w %d - validation failed", ErrWrongRule, intValue)
			}

			if intValue < int64(min) {
				return fmt.Errorf("%d, %w - validation failed", intValue, ErrValueIsLesser)
			}
		case "max":
			max, err := strconv.Atoi(rule.Value)
			if err != nil {
				return fmt.Errorf("%w, %d - validation failed", ErrWrongRule, intValue)
			}

			if intValue > int64(max) {
				return fmt.Errorf("%d, %w - validation failed", intValue, ErrValueIsGreater)
			}
		case "in":
			flag := false
			for _, item := range strings.Split(rule.Value, ",") {
				val, err := strconv.Atoi(item)
				if err != nil {
					return fmt.Errorf("%w %d - validation failed", ErrWrongRule, intValue)
				}

				if int64(val) != intValue {
					flag = true
					break
				}
			}
			if !flag {
				return fmt.Errorf("value %d not found in rule - validation failed", intValue)
			}
		default:
			return fmt.Errorf("%w %d - validation vailed", ErrUnsupportedRule, intValue)
		}
	}
	return nil
}

func Validate(v interface{}) error {
	structToValidate := reflect.TypeOf(v)
	if structToValidate.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	typeValue := reflect.ValueOf(v).Type()
	var validationErrors ValidationErrors
	for i := 0; i < structToValidate.NumField(); i++ {
		field := structToValidate.Field(i)
		tag := field.Tag.Get(tagValidateName)

		// skip fields without "validate" tag.
		if len(tag) == 0 {
			continue
		}

		rules, err := parsingRules(tag, rulesSplitter, ruleNameValueSplitter)
		if err != nil {
			return err
		}

		value := reflect.ValueOf(v).Field(i)
		fieldValue := typeValue.Field(i)
		switch value.Kind() { //nolint:exhaustive
		case reflect.String:
			if err := stringTypeConstraint(rules, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldValue.Name,
					Err:   fmt.Errorf("validation failed - %w", err),
				})
			}
		case reflect.Int:
			if err := validateMinMax(rules, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldValue.Name,
					Err:   fmt.Errorf("validation failed - %w", err),
				})
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
