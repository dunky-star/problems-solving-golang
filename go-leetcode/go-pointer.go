package main

import (
	"fmt"
	"log"
	"sort"
)

type User struct {
	firstName string
	lastName  string
	email     string
	age       int
}

func main() {
	var myString string = "green"

	var mySlice []any

	myMap := make(map[string]*User)

	user1 := User{
		firstName: "Geoffrey",
		lastName:  "Duncan",
		email:     "dunkystar@example.com",
		age:       35,
	}
	user2 := User{
		firstName: "John",
		lastName:  "Doe",
		email:     "john.doe@example.com",
		age:       39,
	}

	myMap["user1"] = &user1
	myMap["user2"] = &user2
	changeUser(&user1)
	changeUser(&user2)

	for k, v := range myMap {
		log.Printf("Key: %s, Value: %v", k, *v)
	}

	log.Printf("myString is set to %s", myString)
	changeUsingPointer(&myString)
	log.Println("After func call, myString is set to", myString)

	mySlice = append(mySlice, "Geo", "Kaligs", 42, true, user1)
	log.Println("mySlice is set to", mySlice)
	changeSlice(&mySlice)
	log.Println("After func call, mySlice is set to", mySlice)

	sortSlice(mySlice)
	log.Println("After sorting, mySlice is set to", mySlice)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Purple"
	*s = newValue
}

func changeUser(u *User) {
	log.Println("u is set to", u)
	newFirstName := "Jane"
	newLastName := "Doe"
	newEmail := "jane.doe@example.com"
	newAge := 30
	u.firstName = newFirstName
	u.lastName = newLastName
	u.email = newEmail
	u.age = newAge
}

func changeSlice(s *[]any) {
	log.Println("s is set to", s)
	*s = append(*s, "NewItem", 100, false)
}

func sortSlice(slice []any) {
	sort.Slice(slice, func(i, j int) bool {
		return fmt.Sprintf("%v", slice[i]) < fmt.Sprintf("%v", slice[j])
	})
}
