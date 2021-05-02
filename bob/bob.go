// Package bob provides a solution to the Exercism problem on Go track.
package bob

import "strings"

// Hey returns an appropriate reply from Bob on the given remark.
func Hey(remark string) string {
	switch r := strings.TrimSpace(remark); {
	case r == "":
		return "Fine. Be that way!"
	case !isQuestion(r) && isCapitalized(r):
		return "Whoa, chill out!"
	case isQuestion(r) && !isCapitalized(r):
		return "Sure."
	case isQuestion(r) && isCapitalized(r):
		return "Calm down, I know what I'm doing!"
	default:
		return "Whatever."
	}
}

func isCapitalized(remark string) bool {
	// Keep only a-zA-Z characters
	alpha := func(r rune) rune {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			return r
		}
		return '.'
	}
	// Replace every other character with ""
	modified := strings.Map(alpha, remark)
	modified = strings.ReplaceAll(modified, ".", "")
	if modified == "" {
		return false
	}
	// Check if contains only upper case letters
	for _, ch := range modified {
		if ch >= 'a' && ch <= 'z' {
			return false
		}
	}
	return true
}

func isQuestion(remark string) bool {
	return remark[len(remark)-1:] == "?"
}
