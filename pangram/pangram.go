// Package pangram implements a solution to the problem on the Go track of Exercism
package pangram

import "strings"

// IsPangram determines if a sentence is a pangram
func IsPangram(str string) bool {
	if str == "" {
		return false
	}
	var letters = map[rune]uint8{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}

	for _, r := range strings.ToLower(str) {
		if _, ok := letters[r]; ok {
			letters[r]++
		}
	}
	for _, v := range letters {
		if v == 0 {
			return false
		}
	}

	return true
}
