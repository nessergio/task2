package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"task2/models/posts"
	"task2/tests/unit"
	"testing"
)

func TestUpdatePost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupAppRouter(&TestConfig)
	statusCode, bodyString := unit.FakeRequest(t, "PUT", "/api/v1/posts/60", router, map[string]any{
		"author": "Homer",
	})
	assert.Equal(t, http.StatusOK, statusCode)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      60,
		Author:  "Homer",
		Title:   "Title 60",
		Content: "Labore quiquia consectetur neque ut etincidunt. Magnam consectetur dolor consectetur. Neque porro eius numquam labore. Ut neque ut velit quaerat. Quisquam consectetur adipisci aliquam magnam modi est labore. Numquam ut quisquam voluptatem quiquia. Quisquam ipsum tempora ut porro eius dolor. Etincidunt est dolorem non.",
	})), bodyString)
	// Test second update of post and incrementing id
	statusCode, bodyString = unit.FakeRequest(t, "PUT", "/api/v1/posts/60", router, map[string]any{
		"title":   "New Title",
		"content": "Lorem ipsum 2",
	})
	assert.Equal(t, http.StatusOK, statusCode)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      60,
		Author:  "Homer",
		Title:   "New Title",
		Content: "Lorem ipsum 2",
	})), bodyString)
	// Test with malformed author
	statusCode, _ = unit.FakeRequest(t, "PUT", "/api/v1/posts/70", router, map[string]any{
		"author": "New post ??",
	})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	// Test with malformed id
	statusCode, _ = unit.FakeRequest(t, "PUT", "/api/v1/posts/aa", router, map[string]any{
		"content": "Lorem ipsum",
	})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	// Test with large id
	statusCode, _ = unit.FakeRequest(t, "PUT", "/api/v1/posts/170", router, map[string]any{
		"author": "John Doe",
		"title":  "New post 1",
	})
	assert.Equal(t, http.StatusNotFound, statusCode)
}
