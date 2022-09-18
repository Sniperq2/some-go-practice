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

	if len(line) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(line); err == nil {
		return "", ErrInvalidString
	}

	for _, value := range line {
		if unicode.IsDigit(value) {
			if item, err := strconv.Atoi(string(value)); err == nil {
				if item-1 > 0 {
					sb.WriteString(strings.Repeat(stack, item-1))
				} else if item-1 < 0 {
					str := sb.String()
					sb.Reset()
					sb.WriteString(str[:len(str)-1])
				} else {
					sb.WriteString(stack)
				}
			}
		} else {
			stack = string(value)
			sb.WriteString(string(value))
		}
	}

	return sb.String(), nil
}
