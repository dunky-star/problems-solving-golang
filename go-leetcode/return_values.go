package main

import "fmt"

func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l / r, true
}

// Generics

func cloneSlice(s []float64) []float64 {
	result := make([]float64, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func cloneMap(m map[string]float64) map[string]float64 {
	result := make(map[string]float64, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
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

type Reader interface {
	Read([]byte) (int, error)
}

type File struct {
	name string
	size int64
}

func (f File) Read(b []byte) (int, error) {
	// Dummy implementation
	return 0, nil
}

type TCPConn struct {
	address string
	port    int
}

func (t TCPConn) Read(b []byte) (int, error) {
	// Dummy implementation
	return 0, nil
}

func main() {
	// var f File
	// var t TCPConn
	// var r Reader
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

	testScoresSlice := []float64{
		98.5,
		87.3,
		76.3,
		89.0,
		27,
	}

	c1 := cloneSlice(testScoresSlice)
	fmt.Printf("Original slice address: %p, Cloned slice address: %p, Cloned slice: %v\n",
		&testScoresSlice[0], &c1[0], c1)

	testScoresMap := map[string]float64{
		"Math":    98.5,
		"Science": 76.3,
		"English": 89.0,
	}
	c2 := cloneMap(testScoresMap)
	fmt.Printf("\nCloned map: %#v\n", c2)

	// Transient Polymorphism with Generics
	// r = f
	// r = t

}
