// Package hamming provides a solution to the Exercism task Hamming
package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	// Return 0 if inputs are identical, or the number of differences otherwise
	num := 0
	runeA := []rune(a)
	runeB := []rune(b)
	// Return an error if inputs are incomparable
	if len(runeA) != len(runeB) {
		return 0, errors.New("please provide correct inputs")
	}
	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			num++
		}
	}
	return num, nil
}
