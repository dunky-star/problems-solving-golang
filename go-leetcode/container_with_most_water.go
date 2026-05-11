package main

import "math"

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		h := min(height[left], height[right])
		width := right - left
		area := width * h
		maxArea = int(math.Max(float64(maxArea), float64(area)))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return maxArea
}

func main() {
	input := []int{1, 8, 6, 2, 5, 7}
	expected := 28

	result := maxArea(input)
	if result == expected {
		println("Test passed!")
	} else {
		println("Test failed. Expected:", expected, "Got:", result)
	}
}
