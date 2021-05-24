// file: main.go

package main

import "github.com/gin-gonic/gin"

func main() {
	setupServer().Run()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pingEndpoint)

	return r
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
