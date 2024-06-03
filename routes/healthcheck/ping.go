package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingHandler returns "pong" with status code 200
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
