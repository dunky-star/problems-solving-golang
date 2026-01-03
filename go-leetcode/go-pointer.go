package main

import "log"

func main() {
	var myString string = "green"

	log.Printf("myString is set to %s", myString)
	changeUsingPointer(&myString)
	log.Println("After func call, myString is set to", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Purple"
	*s = newValue
}
