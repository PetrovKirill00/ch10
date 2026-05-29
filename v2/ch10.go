// A simple library for adding two numbers
package ch10

import (
	"golang.org/x/exp/constraints"
)

// Add adds two numbers together
//
// More info on additon at [MathIsFun].
//
// [MathIsFun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T constraints.Integer | constraints.Float] (a, b T) T {
	return a + b
}
