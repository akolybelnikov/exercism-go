// Package luhn implements a solution ot the problem on Go track of Exercism
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid determines whether or not the given number is valid per the Luhn formula.
func Valid(str string) bool {
	// Keep track of the total and the current character position
	var sum int
	// Convert to runes to operate on Unicode characters
	runeS := []rune(strings.ReplaceAll(str, " ", ""))
	// Return early if not enough characters
	if len(runeS) <= 1 {
		return false
	}
	even := true
	// Iterate over the runes from end to start
	for i := len(runeS) - 1; i >= 0; i-- {
		// Return false if a character is not numeric
		if !unicode.IsDigit(runeS[i]) {
			return false
		}
		// Double the character numeric value if it's in an odd position
		if even {
			if d, err := strconv.Atoi(string(runeS[i])); err == nil {
				sum += d
			}
		} else {
			if d, err := strconv.Atoi(string(runeS[i])); err == nil {
				dd := d * 2
				if dd > 9 {
					dd -= 9
				}
				sum += dd
			}
		}
		even = !even
	}
	// check for the Luhn's condition and return
	return sum%10 == 0
}
