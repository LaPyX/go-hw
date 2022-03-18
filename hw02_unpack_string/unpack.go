package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	chars := []rune(str)
	var newStr []rune
	length := len(chars)

	for i := 0; i < length; i++ {
		char := chars[i]

		if i == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		nextChar := rune(-1)
		if length > (i + 1) {
			nextChar = chars[i+1]
		}

		if unicode.IsDigit(nextChar) && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		if !unicode.IsDigit(nextChar) && !unicode.IsDigit(char) {
			newStr = append(newStr, char)

			continue
		}

		charInt, _ := strconv.Atoi(string(nextChar))

		if unicode.IsDigit(char) || charInt == 0 {
			continue
		}

		charRepeat := strings.Repeat(string(char), charInt)
		charRune := []rune(charRepeat)

		newStr = append(newStr, charRune...)
	}

	return string(newStr), nil
}
