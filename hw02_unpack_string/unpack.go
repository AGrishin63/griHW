package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	s "strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	// Place your code here
	var newStr string = ""
	var digitCount int = 1
	digitCount++
	for i := 0; i < len(input); i++ {
		var ch rune = rune(input[i])
		if unicode.IsDigit(ch) {
			digitCount++
			if digitCount > 1 {
				return "", ErrInvalidString
			} else {
				n, _ := strconv.Atoi(string(input[i]))
				newStr += s.Repeat(string(input[i-1]), n-1)
			}
		} else {
			digitCount = 0
			newStr += string(input[i])
		}

		//fmt.Printf("%#U   %s \n", i, input[i])
		//fmt.Printf("%#U  %v \n", i, ch)
		//}
		//switch expression {
		//case condition:

	}
	return newStr, nil
}
