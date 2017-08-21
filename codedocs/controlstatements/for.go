package control

import "fmt"

func forloop() {
	/*	Go has only one looping construct, the for loop.
		The basic for loop has three components separated by semicolons,
			the init statement: executed before the first iteration
			the condition expression: evaluated before every iteration
			the post statement: executed at the end of every iteration
	*/
	var sum int

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Unlike other languages like C, Java, or Javascript there are no
	// parentheses surrounding the three components of the for statement
	// and the braces { } are always required

	// The init and post statement are optional.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Infinite loop
	for {
	}
}
