package main

import (
	"bufio"
	"fmt"
	"os"
	"reviews/menu"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop1: // Label for breaking the loop in a switch statment
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.AddItem()

		case "q":
			break loop1
		default:
			fmt.Println("Unknown option")
		}
	}

}
