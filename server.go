package main

import (
	"fmt"

	// gin <3
	"github.com/gin-gonic/gin"

	// routes
	"gin-prototype/controllers/index"


)

func main() {
	fmt.Print("Starting...")

	// init
	router := gin.New()
	// default middleware
	router = gin.Default()

	// controllers
	index.Init(router)


	// catch all request to non-existing endpoints
	router.NoRoute(func(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": fmt.Sprintf("Bad Request - %s %s is not a valid endpoint!", c.Request.Method, c.Request.RequestURI),
		})
	})

	router.Run(":")
}

