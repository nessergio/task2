package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler is used in middleware for formatting errors
func ErrorHandler(c *gin.Context) {
	// All chain must proceed first
	c.Next()
	if len(c.Errors) > 0 {
		// Outputting only last error in response, do not show user detailed error
		// -1 to skip setting of the status code (use one from AbortWithError)
		c.JSON(-1, gin.H{"error": c.Errors.Last().Error()})
	}
}

// PanicHandler is custom recovery handler for panic cases
func PanicHandler(c *gin.Context, e any) {
	_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%+v", e))
}
