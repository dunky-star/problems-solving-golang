package main

import (
	"fmt"
	"strconv"
)

func addTwoBinaries(a, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry == 1 {

		if i >= 0 {
			carry += int(a[i]) - '0' // convert char to int
			i--
		}

		if j >= 0 {
			carry += int(b[j]) - '0'
			j--
		}

		result = strconv.Itoa(carry%2) + result // convert int to string and prepend to result
		carry /= 2
	}
	return result
}

func main() {
	a := "1111"
	b := "1101"

	expected := "11100"

	result := addTwoBinaries(a, b)

	fmt.Println("Input A: ", a)
	fmt.Println("Input B: ", b)
	fmt.Println("Expected Output: ", expected)
	fmt.Println("Actual Output: ", result)

	if result == expected {
		fmt.Println("Test Passed!")
	} else {
		fmt.Println("Test Failed!")
	}

}
