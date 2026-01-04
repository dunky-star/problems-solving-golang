package main

import "log"

type User struct {
	firstName string
	lastName  string
	email     string
	age       int
}

func main() {
	var myString string = "green"

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
