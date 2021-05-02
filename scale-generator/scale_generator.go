// Package scale provides a solution to the problem on Go track of Exercism
package scale

import (
	"errors"
	"strings"
)

type chScale int

const (
	chSharp chScale = iota
	chFlat
)

var chromatic = map[chScale][]string{
	chSharp: {"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	chFlat:  {"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"},
}

// Scale generates the musical scale starting with the tonic and following the specified interval pattern
func Scale(tonic string, interval string) (harmony []string) {
	harmony = []string{}
	// Find the scale to use
	scale, err := useScale(tonic)
	if err != nil {
		return
	}
	// Find the starting pitch and shift the scale
	index := findIndex(strings.Title(tonic), chromatic[scale])
	if index == -1 {
		return
	}
	shiftScale := append(chromatic[scale][index:], chromatic[scale][:index]...)
	// If interval empty, return the scale
	if len(interval) == 0 {
		return shiftScale
	}
	// Write the first pitch to harmony
	harmony = append(harmony, shiftScale[0])
	// Use interval to find the rest of the pitches
	i := 0
	for _, step := range interval[:len(interval)-1] {
		switch step {
		case 'M':
			i += 2
		case 'A':
			i += 3
		default:
			i++
		}
		harmony = append(harmony, shiftScale[i])
	}

	return
}

func useScale(tonic string) (chScale, error) {
	switch tonic {
	case "C", "G", "D", "A", "B", "F#", "e", "b", "f#", "c#", "g#", "d#", "a":
		return chSharp, nil
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return chFlat, nil
	default:
		return -1, errors.New("")
	}
}

func findIndex(tonic string, scale []string) int {
	for k, v := range scale {
		if tonic == v {
			return k
		}
	}
	return -1
}
