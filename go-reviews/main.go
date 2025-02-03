package main

import "fmt"

func main() {
	// Constant
	const (
		const1      = 20
		const2      = true
		constPi     = 3.1415 // Flaot64
		constCopied          // Unassigned constant receive previous value
	)

	// String
	var s string = "Foo"
	fmt.Println("String data: ", s)

	// Inttegers
	var num1 int = 99
	fmt.Println("Integer data: ", num1)

	// Float type
	d := 200.233
	var e int = int(d)
	fmt.Println("Type conversion: ", e)

	// Multiple variabe declaration in 1 line
	a, b := 10, 4                 // Go allows multiple variables to be initialized at once
	fmt.Println("Value c: ", a+b) // addition = 15
	fmt.Println("Modulus a(mod)b", a%const1)
}
