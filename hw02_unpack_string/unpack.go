package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev rune
	str := []rune(s)
	for i, char := range str {
		if i == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(char) && unicode.IsDigit(str[i-1]) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(char) {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			if digit == 0 {
				prev = 0
				continue
			}
			result.WriteString(strings.Repeat(string(prev), digit-1))
		} else {
			if prev != 0 {
				result.WriteString(string(prev))
			}
			prev = char
		}
		if i == len(str)-1 && !unicode.IsDigit(char) {
			result.WriteString(string(char))
		}
	}

	return result.String(), nil
}
