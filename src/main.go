package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	// Hides all debug info and warning messages
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/status/:statusCode", func(c *gin.Context) {
		statusCode, err := strconv.Atoi(c.Param("statusCode"))
		if err != nil {
			c.String(500, "Actual error, invalid status code passed")
		} else {
			c.String(statusCode, "This really was successful, but faking status code")
		}
	})

	portNumber := getEnv("PORT", "8080")
	fmt.Printf("Listening on port %s (Use PORT environment variable to change)\n", portNumber)

	r.Run()
}
