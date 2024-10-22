package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/std", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using std http.Handler"))
	})

	// Adding middleware manually
	http.ListenAndServe(":8080", middleware(http.DefaultServeMux))
}

// Sample middleware function
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Custom middleware logic
		next.ServeHTTP(w, r)
	})
}
