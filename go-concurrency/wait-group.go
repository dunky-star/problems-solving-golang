package main

import (
	"fmt"
	"sync"
)

func mainWG() {
	var wg sync.WaitGroup
	wg.Add(1)   // Signals and add that a Gorotuine function is beginning to execute.
	go func() { // Goroutine.
		fmt.Println("This happens asynchronously")
		wg.Done() // Decrement counter by 1.
	}() // Anonymous function -> Invoking the function immediately as soon as we create it.

	fmt.Println("This is synchronous")
	wg.Wait() // Wait till counter is 0.
}
