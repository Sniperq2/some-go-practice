package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(line string) (string, error) {
	var stack string
	var sb strings.Builder
	var isPrevDigit bool

	if len(line) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(line); err == nil {
		return "", ErrInvalidString
	}

	for idx, value := range line {
		if unicode.IsDigit(value) && idx == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(value) && isPrevDigit {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(value) {
			isPrevDigit = true
			item, err := strconv.Atoi(string(value))
			if err != nil {
				return "", ErrInvalidString
			}

			if item-1 > 0 {
				sb.WriteString(strings.Repeat(stack, item-1))
			} else if item-1 < 0 {
				str := sb.String()
				sb.Reset()
				sb.WriteString(str[:len(str)-1])
			}
		} else {
			isPrevDigit = false
			stack = string(value)
			sb.WriteString(string(value))
		}
	}

	return sb.String(), nil
}
