package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}

func main() {
	var p printer
	p = user{username: "Geoffrey", id: 01}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee",
		prices: map[string]float64{"small": 1.66,
			"medium": 1.88,
			"large":  1.90,
		},
	}

	fmt.Println(p.Print())

}

// Type
type myInt int

// Methods indicate a tighter coupling between a function and a type.
func (i myInt) isEven() bool { // Method receiver and doesn't have to be a struct.
	return int(i)%2 == 0
}
