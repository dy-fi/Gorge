package index

import (
	"github.com/gin-gonic/gin"
)

type index struct {}

// Init routes
func Init(e *gin.Engine) {
	c := index{}

	// routes
	e.GET("/", c.root)
}

// serve root
func (*index) root (c *gin.Context) {
	// server html
}