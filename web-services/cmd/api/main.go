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
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"readinglist.dunky.io/internal/data"
	"readinglist.dunky.io/pkg/database"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

var startTime = time.Now()

func main() {
	var cfg config
	godotenv.Load(".env")
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("DB_DSN"), "CockroachDB connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	if cfg.dsn == "" {
		logger.Fatal("db-dsn flag or DB_DSN environment variable must be set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := database.Init(ctx, cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}
	defer database.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(dbPool),
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
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
