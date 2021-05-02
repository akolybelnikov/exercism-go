// Package grains implements a solution to the problem on the Go track of Exercism
package grains

import (
	"errors"
)

// Square calculates how many grains were on a given square.
func Square(input int) (uint64, error) {
	if 0 >= input || input > 64 {
		return 0, errors.New("invalid input")
	}
	return 1 << (input - 1), nil
}

// Total calculate the number of grains of wheat on a chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
