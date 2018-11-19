package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-prototype/controllers"
)

func main() {
	fmt.Print("Starting...")

	router := gin.Default()

	router.GET("/",)
}

