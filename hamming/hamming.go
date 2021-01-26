// Package hamming implements a hamming distance function
package hamming

import (
	"fmt"
	"unicode/utf8"
)

//Distance returns the hamming distance between two strings if they are same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("trings of unequal length string 1 length %d string 2 length %d", len(a), len(b))
	}
	distance := 0
	for i, w := 0, 0; i < len(a); i += w {
		rv1, w := utf8.DecodeRuneInString(a[i:])
		rv2, w := utf8.DecodeRuneInString(b[i:])
		if rv1 != rv2 {
			distance++
		}
		if rv1 == utf8.RuneError || rv2 == utf8.RuneError {
			return 0, fmt.Errorf("rune error")
		}
		i += w
	}
	return distance, nil
}
