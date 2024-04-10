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
	var lastChar rune
	var repeatCount int
	str := []rune(s)

	for i, char := range str {
		if unicode.IsDigit(char) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) && unicode.IsDigit(str[i-1]) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(char) {
			if lastChar == 0 {
				return "", ErrInvalidString
			}
			digit, _ := strconv.Atoi(string(char))
			repeatCount += digit
		} else {
			if repeatCount > 0 {
				result.WriteString(strings.Repeat(string(lastChar), repeatCount-1))
				repeatCount = 0
			}
			if unicode.IsDigit(char) && i > 0 && unicode.IsLetter(str[i-1]) {
				continue
			}
			lastChar = char
			result.WriteString(string(lastChar))
		}
	}

	if repeatCount > 0 {
		result.WriteString(strings.Repeat(string(lastChar), repeatCount))
	}

	return result.String(), nil
}
