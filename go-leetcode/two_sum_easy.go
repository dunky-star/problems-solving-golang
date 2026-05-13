package main

import "fmt"

func twoSum(num []int, target int) []int {
	left := 0
	right := len(num) - 1

	for left < right {
		sum := num[left] + num[right]

		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 11, 7, 15}
	target := 26
	expected := []int{1, 3}

	result := twoSum(numbers, target)

	fmt.Println("Input:    ", numbers)
	fmt.Println("Target:   ", target)
	fmt.Println("Expected: ", expected)
	fmt.Println("Result:   ", result)

	if len(result) == 2 && result[0] == expected[0] && result[1] == expected[1] {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
