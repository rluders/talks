package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Using std http.Handler in Echo
	e.GET("/std", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using std http.Handler"))
	})))

	// Using middleware in Echo
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Custom middleware logic
			return next(c)
		}
	}))

	e.Start(":8080")
}
