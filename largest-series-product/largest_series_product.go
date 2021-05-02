// Package lsproduct implements a solution to the problem on the Go track of Exercism
package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n in a given string
func LargestSeriesProduct(digits string, span int) (int, error) {
	var lsp int
	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	if len(digits) == 0 {
		return 1, nil
	}
	runeS := []rune(digits)
	for idx := 0; idx <= len(runeS)-span; idx++ {
		p, err := findProduct(runeS[idx : idx+span])
		if err != nil {
			return -1, err
		}
		if p > lsp {
			lsp = p
		}

	}
	return lsp, nil
}

func findProduct(digits []rune) (int, error) {
	p := 1
	for _, r := range digits {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return -1, errors.New("digits input must only contain digits")
		}
		p *= d
	}
	return p, nil
}
