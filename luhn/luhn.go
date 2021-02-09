package luhn

import (
	"strings"
	"unicode"
)

//Valid tests if a number is a valid Luhn number
func Valid(s string) bool {
	stripped := strings.ReplaceAll(s, " ", "")
	length := len(stripped)
	double := length%2 == 0
	if length < 2 {
		return false
	}
	sum := 0
	for _, c := range stripped {
		if !(unicode.IsNumber(c)) {
			return false
		}
		n := int(c) - 48
		if double {
			n *= 2
		}
		if n > 9 {
			n -= 9
		}
		sum += n
		double = !double
	}
	if sum%10 != 0 {
		return false
	}
	return true
}
