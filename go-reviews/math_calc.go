package main

import "errors"

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multi(a, b int) int {
	return a * b
}

func Divide(a, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("error: divide by zero error")
	}
	return float32(a / b), nil
}

func Discount(a, b int) float32 {
	return float32(a*b) / 100
}
