package main

import "fmt"

func maxNumber(arr [][]float64) {
	maxNum := arr[0][0]

	for i := range arr {
		for j := range arr[i] {

			if arr[i][j] > maxNum {
				maxNum = arr[i][j]
			}
		}
	}
	fmt.Println("The maximum number from array is: ", maxNum)
}

// func main() {
// 	// Example 2D array
// 	testArr := [][]float64{
// 		{1.1, 2.2},
// 		{7.7, 9.9},
// 	}

// 	maxNumber(testArr)

// }
