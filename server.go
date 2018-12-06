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
	r := gin.New()
	// default middleware
	r = gin.Default()

	// controllers
	index.Init(r)


	// catch all request to non-existing endpoints
	r.NoRoute(func(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": fmt.Sprintf("Bad Request - %s %s is not a valid endpoint!", c.Request.Method, c.Request.RequestURI),
		})
	})

	r.Run()
}

