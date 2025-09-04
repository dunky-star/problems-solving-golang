package main

import "fmt"

func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l / r, true
}

func main() {
	result, ok := divide(10, 2)
	if ok {
		fmt.Println("Result:", result)
	}
}
