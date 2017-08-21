package main

func typeconversions() {

	// The expression T(v) converts the value v to the type T
	i := 42
	f := float64(i)
	u := uint(f)

	// Unlike in C, in Go assignment between items of different type
	// requires an explicit conversion. Try removing the float64 or uint
	// conversions in the example and see what happens.

}
