// Package proverb provides a solution to the exercise on Exercism's Go track.
package proverb

import "fmt"

// Proverb generates the relevant proverb given a list of string inputs.
func Proverb(rhyme []string) (proverb []string) {
	proverb = make([]string, 0, len(rhyme))
	if len(rhyme) > 0 {
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	if len(rhyme) > 1 {
		lines := make([]string, 0, len(rhyme)-1)
		iter := zipSlices(rhyme[:len(rhyme)-1], rhyme[1:])
		for tuple := iter(); tuple != nil; tuple = iter() {
			lines = append(lines, fmt.Sprintf("For want of a %s the %s was lost.", tuple[0], tuple[1]))
		}
		proverb = append(lines, proverb...)
	}

	return proverb
}

// A function iterator zips slices into tuple-like lists
func zipSlices(lists ...[]string) func() []string {
	zip := make([]string, len(lists))
	i := 0
	return func() []string {
		for j := range lists {
			if i >= len(lists[j]) {
				return nil
			}
			zip[j] = lists[j][i]
		}
		i++
		return zip
	}
}
