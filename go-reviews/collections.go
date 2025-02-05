package main

import (
	"fmt"
	"slices"
)

func main() {
	// Array
	var arr [3]int
	fmt.Println("Array: ", arr)
	// SLices - is what we called Refenrence data tyoe.
	var s = []int{1, 2, 3}
	s = append(s, 5, 10, 15)   // Add elements to the slice
	s = slices.Delete(s, 0, 1) // Remove element at index 0 from the slice
	fmt.Println("Slice: ", s)
	// Map
	var map1 = map[string]int{"foo": 10, "bar": 20, "dunky": 500, "quincy": 200} // Map literal
	map1["bar"] = 99                                                             // Update value in the map
	delete(map1, "quincy")                                                       // Delete/Remove entry from a map
	fmt.Println(map1)
	//Structs
}
