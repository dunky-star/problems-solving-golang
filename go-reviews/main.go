package main

import "fmt"

func main() {
	// Constant
	const (
		const1      = 20
		const2      = true
		constPi     = 3.1415 // Flaot64
		constCopied          // Unassigned constant receive previous value
		constIota   = iota
	)

	// String
	var s string = "Foo"
	fmt.Println("String data: ", s)

	// Inttegers
	var num1 int = 99
	fmt.Println("Integer data: ", num1)

	// Float type
	f64 := 200.233
	var e int = int(f64)
	fmt.Println("Type conversion: ", e)

	// Multiple variabe declaration in 1 line
	x, y := 10, 4                 // Go allows multiple variables to be initialized at once
	fmt.Println("Value c: ", x+y) // addition = 15
	fmt.Println("Modulus a(mod)b", x%const1)
	fmt.Println("IOTA: ", constIota)

	// Pointers
	a := 100
	b := &a //(b points to a)
	c := *b // Derefence to get value stored
	fmt.Println("[b] points to [a]", b)
	fmt.Println("Dereferencing to extract value from b: ", c)
}
