// Package raindrops has another string function that converts a number
package raindrops

import "strconv"

// Convert converts a number to a string based on some rules
func Convert(n int) string {
	o := ""
	if n%3 == 0 {
		o += "Pling"
	}
	if n%5 == 0 {
		o += "Plang"
	}
	if n%7 == 0 {
		o += "Plong"
	}
	if o == "" {
		return strconv.Itoa(n)
	}
	return o
}
