package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page")
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the book view page")
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the book create page")
}
