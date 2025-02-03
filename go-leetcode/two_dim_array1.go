package main

import "fmt"

func maxNumber(arr [][]float64) {
	maxNum := arr[0][0]

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {

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
