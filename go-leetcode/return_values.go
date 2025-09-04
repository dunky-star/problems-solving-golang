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
	return (i)%2 == 0
}

// Interface
type printer interface {
	Print() string
}

type user struct {
	id       int
	username string
	email    string
}

func (u user) Print() string { // Method receiver of type user
	return fmt.Sprintf("%v (%v) (%v)\n", u.id, u.username, u.email)
}

func (u *user) changeEmail(newEmail string) { // Pointer receiver to modify the original struct
	u.email = newEmail
}

func main() {
	result, ok := divide(10, 2)
	if ok {
		fmt.Println("Result:", result)
	}
	num := myInt(4)
	fmt.Println("Is even:", num.isEven())

	// Using methods on struct
	u := user{id: 1, username: "dunky", email: "dunky@example.com"}
	fmt.Print(u.Print())
	u.changeEmail("dunky@newdomain.com")
	fmt.Print(u.Print())

}
