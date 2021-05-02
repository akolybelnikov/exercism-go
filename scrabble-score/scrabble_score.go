// Package scrabble provides a solution to the problem on the Go track of Exercism
package scrabble

import "strings"

type score struct {
	letters map[rune]int
}

var values = map[rune]int{'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1, 'D': 2, 'G': 2, 'B': 3, 'C': 3, 'M': 3, 'P': 3, 'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4, 'K': 5, 'J': 8, 'X': 8, 'Q': 10, 'Z': 10}

// Score computes the Scrabble score for the given word
func Score(input string) (result int) {
	result = 0
	if input == "" {
		return
	}
	s := score{letters: map[rune]int{}}
	for _, l := range strings.ToUpper(input) {
		s.letters[l]++
	}
	result = s.evaluate()
	return
}

func (s score) evaluate() int {
	var total = 0
	for k, v := range s.letters {
		if val, ok := values[k]; ok {
			total += v * val
		}

	}
	return total
}
