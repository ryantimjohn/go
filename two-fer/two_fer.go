// Package twofer provides a function to share with someone, stringwise
package twofer

// ShareWith returns a string shareing with another person, defaults to you if empty
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
