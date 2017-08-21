package control

import (
	"fmt"
	"math"
)

// basic if
func sqrt(x float64) string {
	// Go's if statements are like its for loops;
	// the expression need not be surrounded by parentheses ( )
	// but the braces { } are required.

	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with short declaration
func pow(x, n, lim float64) float64 {
	// Like for, the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if with else
func pow2(x, n, lim float64) float64 {
	// Variables declared inside an if short statement are also available inside any of the else blocks.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
