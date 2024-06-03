package unit

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AssertMarshallJson(t *testing.T, payload any) []byte {
	jsonValue, err := json.Marshal(payload)
	assert.Equal(t, err, nil)
	return jsonValue
}

func FakeRequest(t *testing.T, method string, url string, router *gin.Engine, payload any) (int, string) {
	w := httptest.NewRecorder()
	var bodyReader io.Reader
	if payload != nil {
		bodyReader = bytes.NewBuffer(AssertMarshallJson(t, payload))
	}
	req, err := http.NewRequest(method, url, bodyReader)
	assert.Equal(t, err, nil)
	router.ServeHTTP(w, req)
	return w.Code, w.Body.String()
}
