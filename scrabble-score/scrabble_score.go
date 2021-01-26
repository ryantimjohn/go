// Package scrabble allows you to get a words scrabble score
package scrabble

import (
	"strings"
)

//Score takes a string and gives you back the scrabble score of it.
func Score(s string) int {
	o := 0
	for _, l := range strings.ToUpper(s) {
		switch l {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			o++
		case 'D', 'G':
			o += 2
		case 'B', 'C', 'M', 'P':
			o += 3
		case 'F', 'H', 'V', 'W', 'Y':
			o += 4
		case 'K':
			o += 5
		case 'J', 'X':
			o += 8
		case 'Q', 'Z':
			o += 10
		}
	}
	return o
}
