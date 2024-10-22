package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Using std http.Handler in Gorilla Mux
	router.HandleFunc("/std", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using std http.Handler"))
	})

	// Using middleware in Gorilla Mux
	router.Use(loggingMiddleware)

	http.ListenAndServe(":8080", router)
}

// Sample logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Custom middleware logic
		next.ServeHTTP(w, r)
	})
}
