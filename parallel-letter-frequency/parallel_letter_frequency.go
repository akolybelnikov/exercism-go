// Package letter implements a solution to the problem on Exercism
package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in each text in a list of texts concurrently
// and returns this data as a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	mapStream := make(chan FreqMap, len(texts))
	var wg sync.WaitGroup
	wg.Add(len(texts))

	for _, t := range texts {
		go func(text string) {
			m := Frequency(text)
			mapStream <- m
			wg.Done()
		}(t)
	}

	go func() {
		wg.Wait()
		close(mapStream)
	}()

	fm := make(FreqMap)
	for m := range mapStream {
		for k, v := range m {
			fm[k] += v
		}
	}
	return fm
}
