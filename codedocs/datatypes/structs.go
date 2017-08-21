package datatypes

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func structs() {
	// -------------------------------- Basics --------------------------------
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4 // accessing struct values

	// -------------------------------- Pointers to structs --------------------------------
	p := &v
	(*p).X = 1e9 // To access the field X of a struct when we have the struct pointer p we could write (*p).X
	p.Y = 123    // However, that notation is cumbersome, so the language permits us instead to write just p.X,
	// without the explicit dereference, so implicit dereference is implemented in Go

	fmt.Println(v)

	// -------------------------------- Struct Literals --------------------------------
	var (
		v1  = Vertex{1, 2}  // has type Vertex
		v2  = Vertex{X: 1}  // Y:0 is implicit
		v3  = Vertex{}      // X:0 and Y:0
		ptr = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, ptr, v2, v3)
}
