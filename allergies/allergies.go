// Package allergies implements a solution to the problem on the Go track of Exercism
package allergies

var substances = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies determines the full list of allergies given the person's score.
func Allergies(score uint) (results []string) {
	for k, v := range substances {
		// calculate the allergy score by shifting the bits
		r := uint(1 << k)
		if r&score == r {
			results = append(results, v)
		}
	}
	return
}

// AllergicTo determines whether or not the person is allergic to a given item given their score.
func AllergicTo(score uint, substance string) bool {
	for k, v := range substances {
		if v == substance {
			// calculate the allergy score by shifting the bits
			r := uint(1 << k)
			if r&score == r {
				return true
			}
		}
	}
	return false
}
