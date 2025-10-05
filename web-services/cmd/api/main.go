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
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

var startTime = time.Now()

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create the HTTP Server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),     // Set the routes for the server
		ReadTimeout:  10 * time.Second, // Maximum duration for reading the entire request, including the body
		WriteTimeout: 30 * time.Second, // Maximum duration before timing out writes of the response
		IdleTimeout:  time.Minute,      // Maximum amount of time to wait for the next request when keep-alives are enabled
	}

	// Define the server address and port
	addr := fmt.Sprintf(":%d", cfg.port)

	logger.Printf("Server is running on http://localhost%s\n", addr)

	// Start the server and log any error if it fails
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
