// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provide by Gin
	router = gin.Default()

	// Process the template at the start so that they don't have to be loaded
	// from the disk again, This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}
