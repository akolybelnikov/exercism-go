// Package anagram implements a solution to the problem on the Go track of Exercism
package anagram

import (
	"strings"
)

// Detect selects a sublist of anagrams of a given word
func Detect(input string, candidates []string)(anagrams []string)  {
	input = strings.ToLower(input)
	for _, c := range candidates {
		cl := strings.ToLower(c)
		if len(c) != len(input) || input == cl {
			continue
		}
		if isAnagram(input, cl) {
			anagrams = append(anagrams, c)
		}
	}
	return
}

func isAnagram(input string, candidate string) bool {
	letters := map[rune]uint8{}
	for _, r := range input {
		letters[r]++
	}
	for _, l := range candidate {
		if _, ok := letters[l]; !ok {
			return false
		} else {
			letters[l]--
		}
	}
	for _, v := range letters{
		if v != 0 {
			return false
		}
	}
	return true
}
