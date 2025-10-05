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

	*** Routing is limited to exact path matching and http.ServeMux does not support path parameters (e.g. /users/{id})
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/greeting", greetingHandler)
	mux.HandleFunc("/v1/healthcheck", healthCheckHandler)

	PORT := ":4000"
	fmt.Printf("Server is running on http://localhost%s\n", PORT)

	// Start the server and log any error if it fails
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println(err)
	}
}

var startTime = time.Now()

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Labor Omnia Vincit — %s\n", time.Now().Format(time.RFC1123))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Calculate uptime dynamically
	uptime := time.Since(startTime).Truncate(time.Second)
	status := map[string]interface{}{
		"status":    "available",
		"uptime":    uptime.String(),
		"timestamp": time.Now().Format(time.RFC3339),
	}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
	}
}
