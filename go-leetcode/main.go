package main

import (
	"fmt"
	"net/http"
)

func mainBkup() {
	// Handle the root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services are easy with Go!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	// Print a message indicating the server is starting
	fmt.Println("Starting web server on port 9090...")

	// Starting the web server
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting the web server: ", err)
	}

}
