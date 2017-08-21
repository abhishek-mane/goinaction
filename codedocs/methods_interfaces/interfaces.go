package main

// A type implements an interface by implementing its methods. 
// There is no explicit declaration of intent, no "implements" keyword.

// Implicit interfaces decouple the definition of an interface from its implementation, 
// which could then appear in any package without prearrangement. 

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v	// ERROR

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


/* Interface Values :-
	Under the covers, interface values can be thought of as a tuple of a value and a concrete type: 
		(value, type)
	 An interface value holds a value of a specific underlying concrete type.
	 Calling a method on an interface value executes the method of the same name on its underlying type. 
*/

/* OUTPUT :-
	(&{Hello}, *main.T)
	Hello
	(3.141592653589793, main.F)
	3.141592653589793
*/