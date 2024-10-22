package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Using std http.Handler in Gin
	router.GET("/std", gin.WrapH(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Using std http.Handler"))
	})))

	// Using middleware in Gin
	router.Use(gin.Logger(), gin.Recovery())

	router.Run(":8080")
}
