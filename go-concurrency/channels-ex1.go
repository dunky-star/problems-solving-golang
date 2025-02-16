package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2) // Add 2 to account for both goroutines
	go func() {
		defer wg.Done()
		ch <- 42

	}()

	go func() {
		defer wg.Done()
		fmt.Println("Receving value from the channel: ", <-ch)

	}()
	wg.Wait()

	// Using select statement in Channel
	ch1, ch2 := make(chan string, 2), make(chan string, 2) // Buffered channels

	// Send messages in goroutines to prevent blocking
	func() {
		ch1 <- "Message to channel 1"
	}()

	func() {
		ch2 <- "Message to channel 2"
	}()

	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("no messages available")
	}
}
