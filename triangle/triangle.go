// Package triangle determines if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral, all 3 sides of the same length
	Iso             // isosceles, at least 2 sides of the same length
	Sca             // scalene, all 3 sides of a different length
)

type Triangle struct {
	sideA, sideB, sideC float64
}

// KindFromSides returns the kind of the triangle given its sides.
func KindFromSides(a, b, c float64) Kind {
	fs := []float64{a, b, c}
	if !validArgs(fs) {
		return NaT
	}
	t := Triangle{a, b, c}
	if !t.isTriangle() {
		return NaT
	}
	return t.ofKind()
}

func validArgs(fs []float64) bool {
	for _, f := range fs {
		if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) || f <= 0 {
			return false
		}
	}
	return true
}

func (triangle *Triangle) isTriangle() bool {
	return triangle.sideA <= triangle.sideB+triangle.sideC &&
		triangle.sideB <= triangle.sideA+triangle.sideC &&
		triangle.sideC <= triangle.sideA+triangle.sideB
}

func (triangle *Triangle) ofKind() Kind {
	if triangle.sideA == triangle.sideB && triangle.sideA == triangle.sideC {
		return Equ
	}
	if triangle.sideA == triangle.sideB ||
		triangle.sideA == triangle.sideC || triangle.sideB == triangle.sideC {
		return Iso
	}
	return Sca
}
