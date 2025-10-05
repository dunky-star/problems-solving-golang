package main

/*
Limitation of net/http package:
-------------------------------
Basic Functionality: The net/http package provides basic HTTP server and client functionality,
but it lacks advanced features found in more comprehensive web frameworks.
For example, it does not have built-in support for routing, middleware, or templating.
	*** You must build your own wrapper functions or use third-party routers such as Gorilla Mux, Chi, or Gin for features like:

	-Request logging
	-Authentication
	-Panic recovery
	-CORS / rate limiting
*/

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", greetingHandler)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)

	// Start the server and log any error if it fails
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Labor Omnia Vincit — %s\n", time.Now().Format(time.RFC1123))
}
