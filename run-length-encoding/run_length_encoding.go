// Package encode implements a solution to the problem on Go track of exercism
package encode

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

// CountCharacters returns the number of the same characters in a subsequence
func CountCharacters(runeS []rune) int {
	count := 1
	for i := 1; i < len(runeS); i++ {
		if runeS[i] == runeS[0] {
			count++
			continue
		} else {
			break
		}
	}
	return count
}

// RunesToInt returns the integer value of a subsequence of one or more digit characters and the number of them.
func RunesToInt(runeS []rune) (int, int) {
	i := 0
	for unicode.IsDigit(runeS[i]) {
		i++
	}
	s := string(runeS[:i])
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return res, i
}

// RunLengthEncode implements run-length encoding.
func RunLengthEncode(s string) string {
	var b bytes.Buffer
	// Convert the string argument to runes
	runeS := []rune(s)
	i := 0
	// Iterate over the slice of runes and count equal characters in each subsequence
	for i < len(runeS) {
		n := CountCharacters(runeS[i:])
		// If the count is larger than one, write the count to the result as string
		if n > 1 {
			b.WriteString(strconv.Itoa(n))
		}
		// Write the character to the result
		b.WriteRune(runeS[i])
		i += n
	}

	return b.String()
}

// RunLengthDecode implements run-length and decoding.
func RunLengthDecode(s string) string {
	var b bytes.Buffer
	// Convert the string argument to runes
	runeS := []rune(s)
	i := 0
	// Iterate over the slice of runes and identify digital characters and their integer value
	for i < len(runeS) {
		// If a next character is a digit, find out the integer value of the subsequence of digits
		if unicode.IsDigit(runeS[i]) {
			n, x := RunesToInt(runeS[i:])
			for j := 0; j < n; j++ {
				// Write the character that follows the subsequence of digits to the result
				// the number of times corresponding to the integer value of the subsequence
				b.WriteRune(runeS[i+x])
			}
			i += x + 1
		} else {
			// If a next character is not a digit, simply write it to the result one time
			b.WriteRune(runeS[i])
			i++
		}
	}

	return b.String()
}
