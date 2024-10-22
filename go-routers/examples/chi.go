package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	// Using std http.Handler in Chi
	router.Get("/std", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using std http.Handler"))
	}))

	// Using middleware in Chi
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	http.ListenAndServe(":8080", router)
}
