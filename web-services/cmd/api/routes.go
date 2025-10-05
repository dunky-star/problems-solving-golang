package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/greeting", app.greetingHandler)
	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)
	mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	mux.HandleFunc("/v1/books", app.getUpdateDeleteBooksHandler)
	return mux
}