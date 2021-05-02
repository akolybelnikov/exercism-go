// Package etl provides a solution to the problem on Go track of Exercism
package etl

import "strings"

// Transform does the the Transform step of an Extract-Transform-Load
// Returns a data map in the new, desired format
func Transform(legacy map[int][]string) map[string]int {
	new := map[string]int{}
	for k, v := range legacy {
		for _, ch := range v {
			new[strings.ToLower(ch)] = k
		}
	}

	return new
}
