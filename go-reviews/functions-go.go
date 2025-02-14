package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\nNames of students:")
	fmt.Println(strings.Repeat("_", 10))
	greet("\nGeoffrey", "Hills", "Kitty\n")

	name, otherName := "Geoffrey", "Other name Opiyo"
	fmt.Println("Name is: ", name)
	fmt.Println("Other name is: ", otherName)
	myFunc(name, &otherName) // Dereferencing
	fmt.Println(name)
	fmt.Println(otherName)

	fmt.Println("The result of the addition operation is: ", add(10, 30))
	// Calling the function
	result, ok := divide(5, 0)
	// handling the response
	if ok {
		fmt.Println("Division result: ", result)
	} else {
		fmt.Println("Error: divide by zero")
	}

}

func greet(names ...string) { // Variadic parameters
	for _, n := range names {
		fmt.Println(n)
	}
}

// Passing values and pointers
// NB: Use pointers only when you want to share memory otherwise use values
func myFunc(name string, otherName *string) {
	name = "New name"
	*otherName = "Other new name"
}

// Returning a single value
func add(a, b int) int {
	return a * b
}

// Returning multiple values
func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l / r, true
}
