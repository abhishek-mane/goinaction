package control

import (
	"fmt"
	"runtime"
	"time"
)

func switchimp() {
	/* ----------------------------------------------------- */
	// You probably knew what switch was going to look like.
	// A case body breaks automatically, unless it ends with a fallthrough statement.

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	/* ----------------------------------------------------- */
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	/* ----------------------------------------------------- */
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	// below example does not call f()
	i := 0
	switch i {
	case 0:
	case f():
	}
}

func f() int {
	return 10
}
