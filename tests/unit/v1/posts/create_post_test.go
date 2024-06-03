package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"task2/models/posts"
	"task2/routes"
	"task2/tests/unit"
	"testing"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupAppRouter(&TestConfig)
	statusCode, bodyString := unit.FakeRequest(t, "POST", "/api/v1/posts", router, posts.Post{
		Author:  "John Doe",
		Title:   "New post 1",
		Content: "Lorem ipsum",
	})
	assert.Equal(t, http.StatusCreated, statusCode)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      101,
		Author:  "John Doe",
		Title:   "New post 1",
		Content: "Lorem ipsum",
	})), bodyString)
	// Test second creation of post and incrementing id
	statusCode, bodyString = unit.FakeRequest(t, "POST", "/api/v1/posts", router, posts.Post{
		Author:  "Jane Doe",
		Title:   "New post 2",
		Content: "Lorem ipsum 2",
	})
	assert.Equal(t, http.StatusCreated, statusCode)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      102,
		Author:  "Jane Doe",
		Title:   "New post 2",
		Content: "Lorem ipsum 2",
	})), bodyString)
	// Test with empty author
	statusCode, _ = unit.FakeRequest(t, "POST", "/api/v1/posts", router, posts.Post{
		Title:   "New post 3",
		Content: "Lorem ipsum 3",
	})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	// Test with empty title
	statusCode, _ = unit.FakeRequest(t, "POST", "/api/v1/posts", router, posts.Post{
		Author:  "John Doe 4",
		Content: "Lorem ipsum 4",
	})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	// Test with empty content
	statusCode, _ = unit.FakeRequest(t, "POST", "/api/v1/posts", router, posts.Post{
		ID:     1,
		Author: "John Doe",
		Title:  "New post 1",
	})
	assert.Equal(t, http.StatusBadRequest, statusCode)
}
