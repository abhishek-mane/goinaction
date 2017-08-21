package main

/* Nil interface values :-
	A nil interface value holds neither value nor concrete type.
	Calling a method on a nil interface is a run-time error because there is no 
	type inside the interface tuple to indicate which concrete method to call.
*/

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/* OUTPUT :->
	(<nil>, <nil>)
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x0 addr=0x20 pc=0x48a591]

	goroutine 1 [running]:
	main.main()
		C:/Users/gs-1324/AppData/Local/Temp/compile15.go:12 +0x41
*/