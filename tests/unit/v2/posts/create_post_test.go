package posts

import (
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"task2/models/posts"
	v2 "task2/routes/api/v2"
	"testing"
)

func TestCreatePost(t *testing.T) {
	_, api := humatest.New(t)
	postsMemory := posts.FromFileToMemoryDs(TestConfig.InitialDataFile)
	v2.NewPostsController(api, "/api/v2/posts", postsMemory)
	gin.SetMode(gin.ReleaseMode)
	resp := api.Post("/api/v2/posts", map[string]string{
		"author":  "John Doe",
		"title":   "New post 1",
		"content": "Lorem ipsum",
	})
	assert.Equal(t, http.StatusOK, resp.Code)
	require.JSONEq(t, "{\"result\":{\"author\":\"John Doe\", \"content\":\"Lorem ipsum\", \"id\":101, \"title\":\"New post 1\"}}", resp.Body.String())
	resp = api.Post("/api/v2/posts", map[string]string{
		"author":  "John Doe",
		"title":   "New post 1",
		"content": "Lorem ipsum",
	})
	assert.Equal(t, http.StatusOK, resp.Code)
	require.JSONEq(t, "{\"result\":{\"author\":\"John Doe\", \"content\":\"Lorem ipsum\", \"id\":102, \"title\":\"New post 1\"}}", resp.Body.String())
	resp = api.Post("/api/v2/posts", map[string]string{
		"title":   "New post 3",
		"content": "Lorem ipsum 3",
	})
	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)

	resp = api.Post("/api/v2/posts", map[string]string{
		"author":  "John Doe 4",
		"content": "Lorem ipsum 4",
	})
	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)

	resp = api.Post("/api/v2/posts", map[string]any{
		"id":     1,
		"author": "John Doe 4",
		"title":  "New post 1",
	})
	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}
