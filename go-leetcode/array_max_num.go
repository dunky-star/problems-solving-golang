package main

import "fmt"

// a non-empty array of positive integers called nums,
// and you wanted to answer the question: "What is the largest number in nums?"

func maxNum(nums []int) int {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main() {

	numsArray := []int{10, 2, 3, 100, 500, 4}
	fmt.Printf("The max number in array %v is: %v", numsArray, maxNum(numsArray))
}
