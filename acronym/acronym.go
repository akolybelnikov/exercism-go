// Package acronym provides a solution to the problem on Go track at Exercism.
package acronym

import (
	"bytes"
	"regexp"
	"strings"
)

// Abbreviate converts the given phrase to its acronym.
func Abbreviate(s string) string {
	if s != "" {
		var result bytes.Buffer
		sanitize := func(r rune) rune {
			if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == '\'' {
				return r
			}
			return ' '
		}
		sanitized := strings.Map(sanitize, s)
		space := regexp.MustCompile(`\s+`)
		sanitized = space.ReplaceAllString(sanitized, " ")
		words := strings.Split(sanitized, " ")
		for _, word := range words {
			result.WriteString(strings.ToUpper(string(word[0])))
		}
		return result.String()
	}

	return ""
}
