package datatypes

import (
	"fmt"
)

func pointers() {
	i, j := 21, 22

	var p *int
	p = &i
	q := &j

	fmt.Println(*p, *q) // "dereferencing" or "indirecting"

	// Unlike C, Go has no pointer arithmetic.
}
