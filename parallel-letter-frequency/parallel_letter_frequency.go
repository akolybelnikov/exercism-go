// Package letter implements a solution to the problem on Exercism
package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type safeFrequency struct {
	mu sync.Mutex
	fm FreqMap
	wg sync.WaitGroup
}

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
	sf := safeFrequency{fm: make(FreqMap)}
	for _, t := range texts {
		sf.wg.Add(1)
		go sf.concFrequency(t)
	}
	sf.wg.Wait()
	return sf.fm
}

func (f *safeFrequency) concFrequency(s string) {
	defer f.wg.Done()
	for _, r := range s {
		f.mu.Lock()
		f.fm[r]++
		f.mu.Unlock()
	}
}
