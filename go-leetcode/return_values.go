package main

import "fmt"

func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l / r, true
}

// Method on custom type
type myInt int

func (i myInt) isEven() bool { // Method with receiver of type myInt
	return i%2 == 0
}

func main() {
	result, ok := divide(10, 2)
	if ok {
		fmt.Println("Result:", result)
	}
	num := myInt(4)
	fmt.Println("Is even:", num.isEven())
}
