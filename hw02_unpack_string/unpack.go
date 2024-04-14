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
	str := []rune(s)

	var prev rune
	for i, char := range str {
		if unicode.IsDigit(char) && i == 0 {
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
				// Пропускаем предыдущий символ
				prev = 0
			} else {
				// Добавляем повторения предыдущего символа
				if prev != 0 {
					result.WriteString(strings.Repeat(string(prev), digit-1))
				}
			}
		} else {
			// Добавляем текущий символ в результат
			if prev != 0 {
				result.WriteString(string(prev))
			}
			prev = char
		}

		// Если это последний символ и он не цифра, добавляем его в результат
		if i == len(str)-1 && !unicode.IsDigit(char) {
			result.WriteString(string(char))
		}
	}

	return result.String(), nil
}
