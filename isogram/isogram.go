// Package isogram has a function to test if a string is an isogram
package isogram

import (
	"strings"
)

//IsIsogram returns true or false depending on if string is isogram
func IsIsogram(s string) bool {
	if len(s) == 0 {
		return true
	}
	counts := make(map[rune]bool)
	upper := strings.ToUpper(s)
	for _, l := range upper {
		if l >= 65 && l <= 90 {
			if _, ok := counts[l]; ok {
				return false
			}
			counts[l] = true
		}
	}
	return true
}
