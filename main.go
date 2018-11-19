package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Starting...")

	router := gin.Default()

	router.GET("/",)
}

func root(c *gin.Context) {
	
}