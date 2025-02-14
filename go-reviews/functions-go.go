package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\nNames of students:")
	fmt.Println(strings.Repeat("_", 10))
	greet("\nGeoffrey", "Hills", "Kitty\n")

}

func greet(names ...string) { // Variadic parameters
	for _, n := range names {
		fmt.Println(n)
	}
}
