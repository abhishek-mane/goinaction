package datatypes

import "fmt"

// Maps are based on hash tables

func MapFundamentals() {
	// ------------------------------- Creating Maps
	var dict1 map[string]int // Map having key of type int & value of type string, dict1 is nil (no memory space allocated)
	dict1["key1"] = 10       // Runtime Error: panic: runtime error: assignment to entry in nil map

	// using make to create
	dict2 := make(map[int]string) // memory allocated

	dict3 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"} // with initialization

	// The map key can be a value from any built-in or struct type as long as
	// the value can be used in an expression with the == operator. whereas map values has no any such restriction

	// Accessing Maps
	dict4 := map[string]string{}
	// Add the Red color code to the map.
	dict4["Red"] = "#da1337"

	// --------------------------------- Checking existance of key in Map
	// Retrieve the value for the key "Blue".
	value1, exists := dict3["Blue"]
	// Did this key exist?
	if exists {
		fmt.Println(value1)
	}

	// Other way
	// Retrieve the value for the key "Blue".
	value2 := dict3["Blue"]
	// Did this key exist?
	if value2 != "" {
		fmt.Println(value2)
	}
	// The other option is to just return the value and test for the zero value
	// to determine if the key exists. This will only work if the zero value is not a valid value for the map

}

func IteratingOverMaps() {
	// Create a map of colors and color hex codes.
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// Delete key/value pair from map
	delete(colors, "Coral")

}

// Passing maps between functions
// Passing a map between two functions doesn’t make a copy of the map.
// In fact, you can pass a map to a function and make changes to the map,
// and the changes will be reflected by all references to the map
