// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith returns a string with the message, given a name.
// However, if the name is missing, it returns the string 'you'
func ShareWith(name string) string {
	str1 := "One for "
	str2 := ", one for me."
	if name == "" {
		return str1 + "you" + str2
	}
	return str1 + name + str2
}
