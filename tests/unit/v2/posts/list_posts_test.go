package posts

import (
	"encoding/json"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"task2/models/posts"
	"task2/routes/api/v2"
	"testing"
)

func TestListPosts(t *testing.T) {
	_, api := humatest.New(t)
	postsMemory := posts.FromFileToMemoryDs(TestConfig.InitialDataFile)
	v2.NewPostsController(api, "/api/v2/posts", postsMemory)
	gin.SetMode(gin.ReleaseMode)
	resp := api.Get("/api/v2/posts")
	assert.Equal(t, http.StatusOK, resp.Code)
	var result any
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(result.(map[string]any)["result"].([]any)), 100)
}
