// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	upper_remark := strings.ToUpper(remark)
	contains_letters := strings.ContainsAny(upper_remark, "ABCEDEFGHIJKLMNOPQRSTUVWXYZ")
	if remark == upper_remark && contains_letters {
		if remark[len(remark)-1:] == "?" {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else if remark == "" {
		return "Fine. Be that way!"
	} else if remark[len(remark)-1:] == "?" {
		return "Sure."
	} else {
		return "Whatever."
	}
}
