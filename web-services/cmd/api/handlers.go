package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Labor Omnia Vincit — %s\n", time.Now().Format(time.RFC1123))
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Labor Omnia Vincit — %s\n", time.Now().Format(time.RFC1123))
}
