package main

import (
	"errors"
	"fmt"
)

// a non-empty array of positive integers called nums,
// and you wanted to answer the question: "What is the largest number in nums?"

func maxNum(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("error: the array is empty")
	}
	maxNum := nums[0] // Initialize with first value
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum, nil
}

func main() {
	// Test cases
	testCases := [][]int{
		{10, 2, 3, 100, 500, 4}, // Normal case
		{1},                     // Single element
		{7, 7, 7, 7},            // Identical elements
		{0, -1, -2, -3},         // Includes negative numbers and zero
		{-10, -2, -3, -100},     // All negative numbers
		{},                      // Empty array (edge case)
	}

	// numsArray := []int{10, 2, 3, 100, 500, 4}
	// fmt.Printf("The max number in array %v is: %v", numsArray, maxNum(numsArray))
	for _, nums := range testCases {
		result, err := maxNum(nums)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Max in %v is: %v\n ", nums, result)
		}
	}
}
