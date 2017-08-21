package methods_interfaces

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.Abs())
}

// You can only declare a method with a receiver whose type is defined in the
// same package as the method. You cannot declare a method with a receiver whose type
// is defined in another package (which includes the built-in types such as int)

// Method has 2 types of receivers
// 1. value receiver
// 2. pointer receiver
