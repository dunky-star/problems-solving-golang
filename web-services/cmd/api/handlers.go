package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"readinglist.dunky.io/internal/data"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Calculate uptime dynamically
	uptime := time.Since(startTime).Truncate(time.Second)
	status := map[string]interface{}{
		"status":    "available",
		"uptime":    uptime.String(),
		"timestamp": time.Now().Format(time.RFC3339),
	}
	fmt.Fprintf(w, "Version: %s\n", version)
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
	}
}

func (app *application) greetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Labor Omnia Vincit — %s\n", time.Now().Format(time.RFC1123))
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Extract `{id}` from the URL path

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt < 1 {
		http.Error(w, `{"error":"invalid book ID"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Example placeholder JSON response
	response := data.Book{
		ID:        idInt,
		Title:     "The Great Adventure",
		Pages:     350,
		Genres:    []string{"Fiction", "Adventure"},
		Rating:    4.5,
		Published: 2021,
		Version:   1,
		CreatedAt: time.Now(),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error":"Failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (app *application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Mock data — replace this with a DB query later
	books := []data.Book{
		{
			ID:        1,
			CreatedAt: time.Now(),
			Title:     "The Darkening of Tristram",
			Published: 1998,
			Pages:     300,
			Genres:    []string{"Fiction", "Thriller"},
			Rating:    4.5,
			Version:   1,
		},
		{
			ID:        2,
			CreatedAt: time.Now(),
			Title:     "The Legacy of Deckard Cain",
			Published: 2007,
			Pages:     432,
			Genres:    []string{"Fiction", "Adventure"},
			Rating:    4.9,
			Version:   1,
		},
	}

	// Encode and send JSON response
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ") // 2-space indentation
	if err := enc.Encode(books); err != nil {
		http.Error(w, `{"error":"Failed to encode books"}`, http.StatusInternalServerError)
	}

}

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure only POST requests reach here
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var book data.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, `{"error": "invalid JSON payload"}`, http.StatusBadRequest)
		return
	}

	// Close request body after reading
	defer r.Body.Close()

	// For now, we just simulate saving to a DB
	fmt.Printf(" Received new book: %+v\n", book)

	// Set response headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated) // 201 Created

	// Respond with confirmation JSON
	response := map[string]interface{}{
		"message": "Book created successfully",
		"book":    book,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {
	// Enforce PUT-only access for safety
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the book ID from the URL path parameter {id}
	id := r.PathValue("id")

	// You’d parse JSON from r.Body and update in the database
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updating book with ID: %s\n", id)
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only DELETE requests
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract book ID from URL path
	id := r.PathValue("id")

	// You'd delete the book from a database here
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleting book with ID: %s\n", id)
}
