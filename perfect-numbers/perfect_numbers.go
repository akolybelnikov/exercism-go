// Package perfect implements a solution to the problem on the Go track of Exercism.
package perfect

import (
	"errors"
	"math"
)

// Classification contains the classification values for natural numbers.
type Classification string

// Constant Classification values.
const (
	ClassificationDeficient = Classification("deficient")
	ClassificationPerfect   = Classification("perfect")
	ClassificationAbundant  = Classification("abundant")
)

// ErrOnlyPositive is returned when the input is not a positive integer.
var ErrOnlyPositive = errors.New("input is not a positive integer")

// Classify determines if a number is perfect, abundant, or deficient
// based on Nicomachus' (60 - 120 CE) classification scheme for natural numbers.
func Classify(input int64) (Classification, error) {
	if input <= 0 {
		return "", ErrOnlyPositive
	}
	sqrtN := int64(math.Sqrt(float64(input)))
	var i, n int64 = 1, 0
	for i <= sqrtN {
		if input != i && input%i == 0 {
			n += i
			pair := input /i
			if pair < input && pair != i {
				n += input / i
			}
		}
		i++
	}

	switch {
	case n < input:
		return ClassificationDeficient, nil
	case n > input:
		return ClassificationAbundant, nil
	default:
		return ClassificationPerfect, nil
	}
}
