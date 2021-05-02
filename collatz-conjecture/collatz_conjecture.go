// Package collatzconjecture provides a solution to the problem on Go track of exercism
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1.
func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("invalid input")
	}
	if input == 1 {
		return 0, nil
	}
	steps := 0
	for input != 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = (3 * input) + 1
		}
		steps++
	}

	return steps, nil
}
