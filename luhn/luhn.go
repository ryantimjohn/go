package luhn

import (
	"strings"
	"unicode"
)

//Valid tests if a number is a valid Luhn number
func Valid(s string) bool {
	stripped := strings.ReplaceAll(s, " ", "")
	length := len(stripped)
	if length < 2 {
		return false
	}
	sum := 0
	for i, c := range stripped {
		if !(unicode.IsNumber(c)) {
			return false
		}
		n := int(c) - 48
		if (length-i)%2 == 0 {
			n *= 2
		}
		if n > 9 {
			n -= 9
		}
		sum += n
	}
	if sum%10 != 0 {
		return false
	}
	return true
}
