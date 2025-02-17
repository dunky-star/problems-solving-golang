package main

import (
	"fmt"
	"strings"
)

func blueAndRed(arr []string) int {
	count := 0
	if len(arr) == 0 { // Edge case: if the string is empty, return 0
		return 0
	}
	for _, str := range arr {
		if strings.ToLower(str) == "blue" || strings.ToLower(str) == "red" {
			count++
		}
	}
	return count
}

func mainRedBlue() {
	arr := []string{"CAT", "blue", "blue", "CAT", "another", "Red", "bird"}
	fmt.Println("The number of red and blue strings are: ", blueAndRed(arr))
}
