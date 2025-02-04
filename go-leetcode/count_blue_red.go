package main

import (
	"fmt"
	"strings"
)

func blueAndRed(arr []string) int {
	count := 0
	for _, str := range arr {
		if strings.ToLower(str) == "blue" || strings.ToLower(str) == "red" {
			count++
		}
	}
	return count
}

func main() {
	arr := []string{"CAT", "blue", "blue", "CAT", "another", "Red", "bird"}
	fmt.Println("The number of red and blue strings are: ", blueAndRed(arr))
}
