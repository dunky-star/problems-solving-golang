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
	var x = []int{1, 2, 3}
	x = append(x, 5, 10, 15)   // Add elements to the slice
	x = slices.Delete(x, 0, 1) // Remove element at index 0 from the slice
	fmt.Println("Slice: ", x)
	// Map
	var map1 = map[string]int{ // Map literal
		"foo":    10,
		"bar":    20,
		"dunky":  500,
		"quincy": 200,
		"kaligs": 666,
	}
	map1["bar"] = 99 // Update value in the map

	delete(map1, "quincy")                                             // Delete/Remove entry from a map
	v, ok := map1["quincy"]                                            // Optional syntax to query a map
	fmt.Println("Optional syntax to query if an item exists: ", v, ok) // for a key that is not present
	fmt.Println(map1)
	m := map1
	m["baz"] = 999 // Add new item to a map
	fmt.Println(m)

	var map2 = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccion"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}
	map2["other"] = []string{"Hot Chocolate"}
	fmt.Println(map2)
	fmt.Println(map2["coffee "])
	// NB: Maps are not comparable - will raise compile time error (m == map1)

	//Structs
	var s struct { // Declare an anonymous struct
		name   string
		salary float64
		id     int
	}

	s.name = "Geoffrey"
	s.id = 001
	s.salary = 120000
	fmt.Printf("Id: %d, Name: %s, Salary: %.0f\n", s.id, s.name, s.salary)

}
