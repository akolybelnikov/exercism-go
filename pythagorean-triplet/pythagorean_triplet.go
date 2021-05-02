// Package pythagorean implements a solution to the problem on the Go track of Exercism
package pythagorean

import (
	"sort"
)

// Triplet holds a Pythagorean triple
type Triplet [3]int

type intPair [2]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (ts []Triplet) {
	c := 0
	m := 2
	for c <= max {
		for n := 1; n < m; n++ {
			p := intPair{n, m}
			sides, isPythagorean := p.getSides()
			c = sides[2]
			if c > max {
				break
			}
			if sides[0] >= min && isPythagorean {
				t := Triplet{sides[0], sides[1], sides[2]}
				ts = append(ts, t)
			}
		}
		m++
	}

	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) (ts []Triplet) {
	// The value of the first element in sorted triplet can be at most p/3
	for i := 1; i < p/3; i++ {
		// The value of the second element can be at most n/2
		for j := i + 1; j < p/2; j++ {
			k := p - i - j
			if (i*i)+(j*j) == k*k {
				t := Triplet{i, j, k}
				ts = append(ts, t)
			}
		}
	}
	return
}

// According to the Euclid's formula, while m > n > 0:
// a = m**2 - n**2, b = 2 * m * n, c = m**2 + n**2.
func (p *intPair) getSides() ([]int, bool) {
	sides := make([]int, 3)
	sides[0] = p[1]*p[1] - p[0]*p[0]
	sides[1] = 2 * p[1] * p[0]
	sides[2] = p[1]*p[1] + p[0]*p[0]
	sort.Ints(sides)
	return sides, sides[0]*sides[0]+sides[1]*sides[1] == sides[2]*sides[2]
}
