// Package wordcount implements a solution to the problem on the Go track of Exercism
package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string)Frequency  {
	r := regexp.MustCompile(`([\w]+['][\w]+)|(\w+)`)
	words := r.FindAllString(input,-1)
	fq := Frequency{}
	for _, word := range words {
		fq[strings.ToLower(word)]++
	}
	return fq
}
