package unit

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"task2/routes/healthcheck"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/healthcheck", healthcheck.PingHandler)

	statusCode, bodyString := FakeRequest(t, "GET", "/healthcheck", router, nil)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "pong", bodyString)
}
