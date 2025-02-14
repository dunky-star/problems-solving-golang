package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainRun() {

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	}

	in := bufio.NewReader(os.Stdin)

loop1: // Label for breaking the loop in a switch statment
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})

		case "q":
			break loop1
		default:
			fmt.Println("Unknown option")
		}
	}

}
