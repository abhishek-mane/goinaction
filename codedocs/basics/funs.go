package main

import "fmt"

// basic function
func add(x int, y int) int {
	return x + y
}

// func can return any number of values
func swap(x, y string) (string, string) { // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return y, x
}

// func can have named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "Naked" returns, should be used only in short functions
}

// variadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Calling variadic function
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
