package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Using std http.Handler in HttpRouter
	router.Handler("GET", "/std", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using std http.Handler"))
	}))

	// Using middleware manually (httprouter doesn't have native middleware support)
	middleware := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Custom middleware logic
			h.ServeHTTP(w, r)
		})
	}

	http.ListenAndServe(":8080", middleware(router))
}
