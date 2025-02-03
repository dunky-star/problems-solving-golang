package main

import "fmt"

func main() {
	// String
	var s string = "Foo"
	fmt.Println("String data: ", s)

	// Inttegers
	var num1 int = 99
	fmt.Println("Ineteger data: ", num1)

	// Float type
	d := 3.1415 // Flaot64
	var e int = int(d)
	fmt.Println("Type conversion: ", e)

	// Multiple variabe declaration in 1 line
	a, b := 10, 5 // Go allows multiple variables to be initialized at once
	c := a + b    // 15 - addition
	fmt.Println("Value c: ", c)

}
