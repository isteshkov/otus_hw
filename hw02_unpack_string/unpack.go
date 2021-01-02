package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here

	runeStr := []rune(s)

	// empty string
	if len(runeStr) == 0 {
		return "", nil
	}

	// first symbol is digit
	if unicode.IsDigit(runeStr[0]) {
		return "", ErrInvalidString
	}

	resultStr := strings.Builder{}

	for i, v := range s {
		if unicode.IsDigit(v) { //nolint
			// if prev symbol is digit then error (two or more digits in string)
			if unicode.IsDigit(runeStr[i-1]) {
				return "", ErrInvalidString
			}
			intV, _ := strconv.Atoi(string(v))
			if intV > 0 {
				resultStr.WriteString(strings.Repeat(string(runeStr[i-1]), intV))
			}
		} else {
			// last letter in string
			if i == len(runeStr)-1 {
				// write previous if it's not a digit
				if i > 0 && !unicode.IsDigit(runeStr[i-1]) {
					resultStr.WriteString(strings.Repeat(string(runeStr[i-1]), 1))
				}
				// write it once
				resultStr.WriteString(strings.Repeat(string(v), 1))
			} else
			// not first letter and previous is not digit write previous once
			if i > 0 && !unicode.IsDigit(runeStr[i-1]) {
				resultStr.WriteString(strings.Repeat(string(runeStr[i-1]), 1))
			}
		}
	}

	return resultStr.String(), nil
}
