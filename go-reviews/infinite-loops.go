package main

import "fmt"

func mainLoop() {
	i := 0

	// Infinite Loop
	// for {
	// 	fmt.Println(i)
	// 	i += 1
	// }
	for i < 3 {
		fmt.Println("Counter i: ", i)
		i += 1
	}
	fmt.Println("Done with Loop till Condition!")

	// Counter-based Loops
	for i := 1; i < 8; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done with counter-based loops!")
}
