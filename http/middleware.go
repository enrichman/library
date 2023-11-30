package http

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		c.JSON(-1, gin.H{"error": err.Error()})
	}
}
