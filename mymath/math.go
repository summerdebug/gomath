// Package mymath provides basic math functions.
package mymath

import (
	"golang.org/x/exp/constraints"
)

// Add returns the [sum] of two integers.
//
// [sum]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

type Number interface {
	constraints.Integer | constraints.Float
}
