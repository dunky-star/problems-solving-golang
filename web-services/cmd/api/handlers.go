package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, `{"error":"missing book ID"}`, http.StatusBadRequest)
		return
	}

	book, err := app.models.Book.Get(id)
	if err != nil {
		app.logger.Printf("get book error: %v", err)
		http.Error(w, `{"error":"book not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, `{"error":"Failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (app *application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Mock data — replace this with a DB query later
	books, err := app.models.Book.GetAll()
	if err != nil {
		app.logger.Printf("list books error: %v", err)
		http.Error(w, `{"error":"Failed to fetch books"}`, http.StatusInternalServerError)
		return
	}

	// Encode and send JSON response
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ") // 2-space indentation
	if err := enc.Encode(books); err != nil {
		http.Error(w, `{"error":"Failed to encode books"}`, http.StatusInternalServerError)
	}

}

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body
	var input data.Book
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "invalid JSON payload"}`, http.StatusBadRequest)
		return
	}

	// Close request body after reading
	defer r.Body.Close()

	app.logger.Printf("creating book: %+v", input)

	created, err := app.models.Book.Insert(r.Context(), &input)
	if err != nil {
		app.logger.Printf("insert book error: %v", err)
		http.Error(w, `{"error": "failed to create book"}`, http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated) // 201 Created

	// Respond with confirmation JSON
	response := map[string]interface{}{
		"message": "Book created successfully",
		"book":    created,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	var input data.Book
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "invalid JSON payload"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	input.ID = id

	if err := app.models.Book.Update(&input); err != nil {
		app.logger.Printf("update book error: %v", err)
		http.Error(w, `{"error": "failed to update book"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(input); err != nil {
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	if err := app.models.Book.Delete(id); err != nil {
		app.logger.Printf("delete book error: %v", err)
		http.Error(w, `{"error": "failed to delete book"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}
