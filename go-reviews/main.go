package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", Handler)
	fmt.Println("Starting web server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting web server: ", err)
	}
}

// Web services
func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("menu.txt")
	io.Copy(w, f)
}

// Command Line application
func mainPlay1() {
	fmt.Println("What would you like me to scream")
	in := bufio.NewReader(os.Stdin) // NewReader decorator
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")

}

// Go concepts
func mainPlay2() {
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
	b := &a       //(b points to a)
	c := *b       // Derefence to get value stored
	d := new(int) // Pointer to Anonymous memory
	fmt.Println("Anonymous memory of d: ", d)
	fmt.Println("[b] points to [a]", b)
	fmt.Println("Dereferencing to extract value from b: ", c)
}
