package datatypes

import (
	"fmt"
)

// Basics of slice
/*
 A slice is a data structure that provides a way for you to work with and manage collections of data.
 Slices are built around the concept of dynamic arrays that can grow and shrink as you see fit.
 They’re flexible in terms of growth because they have their own built-in function called append,
 which can grow a slice quickly with efficiency.
 You can also reduce the size of a slice by slicing out a part of the underlying memory.
 Slices give you all the benefits of indexing, iteration, and garbage collection
 optimizations because the underlying memory is allocated in contiguous blocks
*/
/*
 They’re three field data structures that contain the metadata Go needs to manipulate
 the underlying arrays (see figure). The three fields are a
 1. pointer to the underlying array,
 2. the length or the number of elements the slice has access to, and
 3. the capacity or the number of elements the slice has available for growth

SLICE -
+---------------------------------------------+
| +-----------+  +-----------+  +-----------+ | Slice of integers:
| |    ptr    |  |  length   |  | Capacity  | |  > Length of integers
| +-----+-----+  +-----------+  +-----------+ |	 > Capacity of elemts
+-------|-------------------------------------+
		|
		|
		|	+--------+ +--------+ +--------+ +--------+ +--------+
		+-->|   10   | |   20   | |   30   | |   40   | |   50   |
			+--------+ +--------+ +--------+ +--------+ +--------+


*/
func sliceBasics() {

	// ---------------------- Creating slice from Array
	array := [6]int{2, 3, 5, 7, 11, 13}
	var slice []int = array[1:4]

	// ---------------------- Create a slice of strings using make
	slice1 := make([]string, 5) // length and capacity of 5 elements

	slice2 := make([]int, 3, 5)
	// length of 3 and capacity of 5 elements.
	// When you specify the length and capacity separately,
	// you can create a slice with available capacity in the underlying array that you don’t have access to initially

	slice2[4] = 10 // trying to access index outside of length will give error

	slice3 := make([]int, 5, 3) // Compiler Error: len larger than cap in make([]int)

	// ----------------------- Create a slice of strings using slice literal
	// Contains a length and capacity of 5 elements.
	slice4 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// Contains a length and capacity of 3 elements.
	slice5 := []int{10, 20, 30}
	// When using a slice literal, you can set the initial length and capacity.
	// All you need to do is initialize the index that represents the length and capacity you need.
	// The following syntax will create a slice with a length and capacity of 100 elements
	slice6 := []string{99: ""}

	// **** Remember, if you specify a value inside the [ ] operator,
	// you’re creating an array. If you don’t specify a value, you’re creating a slice

	// ---------------------- nil slice VS empty slice
	// Empty slices
	slice7 := make([]int, 0) // ptr has address pointing to somewhere
	slice8 := []int{}

	// Nil slices
	var slice9 []int // ptr is nil(null)

	// ----------------------- len & cap
	fmt.Println(len(slice1)) // returns lenght of slice
	fmt.Println(cap(slice1)) // returns capacity of slice

	// -------------------------- Accessing elements of Slices
	// same like arrays
	slice10 := []int{10, 20, 30, 40, 50}
	// Change the value of index 1.
	slice10[1] = 25

	// ----------------------------- creating new slices from existing slices
	slice11 := []int{10, 20, 30, 40, 50}
	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	slice12 := slice11[1:3]
	// slice11 & slice12 sharing same underline array so modifications reflets

	/*** formula
	* For slice[i:j] with an underlying array of capacity k
	  Length: j - i
	  Capacity: k - i
	*/

	// 3rd index option
	slice15 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// Slice the third element and restrict the capacity.
	// Contains a length of 1 element and capacity of 2 elements.
	slice16 := slice15[2:3:4] // 1st index is start, 2nd index length limit & 3rd index is capacity limit
	/*** formula
	* For slice[i:j:k]
	 Length: j - i
	 Capacity: k – i
	*/

	// ------------------------------- Some interesting points
	// When there’s no available capacity in the underlying array for a slice,
	// the append function will create a new underlying array,
	// copy the existing values that are being referenced, and assign the new value.
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice13 := []int{10, 20, 30, 40}
	// Append a new value to the slice.
	// Assign the value of 50 to the new element.
	slice14 := append(slice13, 50)
	// After this operation, slice14 is given its own underlying array and
	// the capacity of the array is doubled from its original size
	// The append operation is clever when growing the capacity of the underlying array.
	// Capacity is always doubled when the existing capacity of the slice is under 1,000 elements.
	// Once the number of elements goes over 1,000, the capacity is grown by a factor of 1.25, or 25%.
	// This growth algorithm may change in the language over time

	// ------------------------------------ Append function is variadic
	slice17 := []int{1, 2}
	fmt.Println(append(slice17, 3, 4))
	fmt.Println(append(slice17, slice12...)) // append all elements of slice12 into slice17

	// Prevent compiler errors
	fmt.Println(array, slice, slice1, slice2, slice3, slice4, slice5, slice6, slice7, slice8, slice9, slice10)
}

func iteratingOverSlices() {
	slice1 := []int{10, 20, 30, 40}
	// Iterate over each element and display each value.
	for index, value := range slice1 {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// If you don’t need the index value, you can use the underscore character to discard the value.
	for _, value := range slice1 {
		fmt.Printf("Value: %d\n", value)
	}

	// The keyword range will always start iterating over a slice from the beginning.
	// If you need more control iterating over a slice, you can always use a traditional for loop.
	// Iterate over each element starting at element 3.
	for index := 2; index < len(slice1); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice1[index])
	}
}

func multidimentionalSlices() {
	// Read topic in Book, 'Go in Action'
}

func passingSlicesInFunctions() {
	// slice is object containing 3 fields, 1st is pointer, 2nd length & 3rd capacity,
	// so it is lightweight to be passed to the functions as value.
	// Since the data associated with a slice is contained in the underlying array,
	// there are no problems passing a copy of a slice to any function.
	// Only the slice is being copied, not the underlying array.
}
