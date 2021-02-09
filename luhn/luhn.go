package luhn

import (
	"strings"
	"unicode"
)

//Valid tests if a number is a valid Luhn number
func Valid(s string) bool {
	stripped := strings.ReplaceAll(s, " ", "")
	double := len(stripped)%2 == 0
	valid := len(stripped) > 1
	sum := 0
	for _, c := range stripped {
		if !(unicode.IsDigit(c)) {
			valid = false
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
		if !valid {
			break
		}
	}
	if sum%10 != 0 {
		valid = false
	}
	return valid
}
