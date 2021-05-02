// Package darts provides a solution to the problem on the Go track of Exercism
package darts

import "math"

// Score calculates and returns the earned points in a single toss of a Darts game.
func Score(x, y float64) (score int) {
	// in order to lie within the circle area the distance of the point (x,y) from (0,0) coordinates
	// must be less or equal to the radius of the circle
	r := radius(x, y)
	switch {
	case r <= 1.0:
		score = 10
	case r <= 5.0:
		score = 5
	case r <= 10.0:
		score = 1
	default:
		score = 0
	}
	return
}

func radius(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))
}
