package datatypes

import (
	"fmt"
)

// Array basics
// An array in Go is a fixed-length data type that contains a contiguous block of elements of the same type.
func ArraysImpl() {

	// Once an array is declared, neither the type of data being stored nor its length can be changed
	var arr [5]int

	// When variables in Go are declared, they’re always initialized to their zero value for their respective type, and arrays are no different
	arr1 := [5]int{10, 20, 30, 40, 50}   // size with initialization
	arr2 := [...]int{10, 20, 30, 40, 50} // size will be calculated from right hand side
	arr3 := [5]int{1: 10, 2: 20}         //initialize only required elements in an array

	// Array of pointers
	arr4 := [5]*int{0: new(int), 1: new(int)}
	*arr4[0] = 10
	*arr4[1] = 20

	// An array is a value in Go. This means you can use it in an assignment operation.
	// The variable name denotes the entire array and, therefore, an array can be assigned to other arrays of the same type.
	// Only arrays of the same type can be assigned
	var arr5 [5]string
	arr6 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	arr5 = arr6 // copy all elements of arr6 into arr5

	// compiler error as lenght mismatch
	var arr7 [4]string
	arr8 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	/* ------------------
	arr7 = arr8 // Compiler Error : cannot use array2 (type [5]string) as type [4]string in assignment
	---------------------*/

	// Copying an array of pointers copies the pointer values and not the values that the pointers are pointing to.
	var arr9 [3]*string
	arr10 := [3]*string{new(string), new(string), new(string)}
	*arr10[0] = "Red"
	*arr10[1] = "Blue"
	*arr10[2] = "Green"
	arr9 = arr10

	fmt.Println(arr, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9, arr10)

}

// Multidimentional Arrays
func MultiDimentionalArrays() {
	// Arrays are always one-dimensional, but they can be composed to create multidimensional arrays.
	// Multidimensional arrays come in handy when you need to manage data that may have parent/child relationships
	// or is associated with a coordinate system
	var arr1 [4][2]int
	arr2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	arr3 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	arr4 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	// Accessing
	var arr5 [2][2]int
	arr5[0][0] = 10
	arr5[0][1] = 20
	arr5[0][1] = 20
	arr5[1][0] = 30
	arr5[1][1] = 40

	// Because an array is a value, you can copy individual dimensions
	var arr6 [2]int = arr5[1]

	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)

}

// Passing an arrays between functions
func PassingArraysToFunctions() {
	// Passing an array between functions can be an expensive operation in terms of memory and performance.
	// When you pass variables between functions, they’re always passed by value. When your variable is an array,
	// this means the entire array, regardless of its size, is copied and passed to the function
	// the better way is to pass an pointer to an array

	var arr1 [1e6]int

	func(ptr *[1e6]int) {

		//...

	}(&arr1) // passing an address of array to the function

}
