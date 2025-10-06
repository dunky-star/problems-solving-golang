package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/greeting", app.greetingHandler)
	mux.HandleFunc("GET /v1/healthcheck", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/books/{id}", app.getBookHandler)  // Get a specific book
	mux.HandleFunc("GET /v1/books", app.listBooksHandler)   // GET → list all books
	mux.HandleFunc("POST /v1/books", app.createBookHandler) // POST → create new book
	mux.HandleFunc("PUT /v1/books/{id}", app.updateBookHandler)
	mux.HandleFunc("DELETE /v1/books/{id}", app.deleteBookHandler)
	return mux
}
