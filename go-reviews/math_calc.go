package main

import (
	"errors"
	"math/rand"
	"time"
)

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

func randomNumber() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	min, max := 0, 100
	itemDiscount := r.Intn(max-min+1) + min
	return itemDiscount
}

func RandomNumber() int {
	return randomNumber()
}
