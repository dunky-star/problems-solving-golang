package main

import (
	"errors"
	"fmt"
)

// We consider errors as values
func mainErr() {
	err := errors.New("this is an error")
	fmt.Println(err)

	err2 := fmt.Errorf("this error wraps the first one: %w", err)
	fmt.Println(err2)
}
