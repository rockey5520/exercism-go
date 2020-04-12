// Package triangle has function to determine the type of the triangle.
package triangle

import (
	"math"
	"sort"
)

// Kind is a type.
type Kind int

//
const (
	// NaT represents not a triangle
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	sides := [3]float64{a, b, c}
	sort.Float64s(sides[:])

	for _, s := range sides {
		if s <= 0 || math.IsNaN(s) || math.IsInf(s, 0) {
			return NaT
		}
	}

	if sides[0]+sides[1] < sides[2] {
		return NaT
	} else if sides[0] == sides[1] && sides[0] == sides[2] {
		return Equ
	} else if sides[0] == sides[1] || sides[1] == sides[2] || sides[2] == sides[0] {
		return Iso
	}
	return Sca
}
