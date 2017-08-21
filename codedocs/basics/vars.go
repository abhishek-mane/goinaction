package main

/* 
declarations
Zero Value => Variables declared without an explicit initial value are given their zero value.
The zero value is:
	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
*/
var b1, b2, b3 bool // package level scope
var b4, b6, b7 bool = true, false, true

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var c, python, java = true, false, "no!"

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func test() {
	var i1 int // function level scope

	// short variable declarations
	k := 3
	s1, s2, s3 := 1, false, "no!"

	// variable declarations may be "factored" into blocks, as with import statements.
}

// Short variable declarations not allowed outside function
// Outside a function, every statement begins with a keyword (var, func, and so on)
test := true

// 