package posts

import (
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"task2/models/posts"
	"task2/routes/api/v2"
	"testing"
)

func TestUpdatePost(t *testing.T) {
	_, api := humatest.New(t)
	postsMemory := posts.FromFileToMemoryDs(TestConfig.InitialDataFile)
	v2.NewPostsController(api, "/api/v2/posts", postsMemory)
	gin.SetMode(gin.ReleaseMode)
	resp := api.Put("/api/v2/posts/60", map[string]string{
		"author": "Homer",
	})
	assert.Equal(t, http.StatusOK, resp.Code)
	require.JSONEq(t, "{\"result\":{\"author\":\"Homer\", \"content\":\"Labore quiquia consectetur neque ut etincidunt. Magnam consectetur dolor consectetur. Neque porro eius numquam labore. Ut neque ut velit quaerat. Quisquam consectetur adipisci aliquam magnam modi est labore. Numquam ut quisquam voluptatem quiquia. Quisquam ipsum tempora ut porro eius dolor. Etincidunt est dolorem non.\", \"id\":60, \"title\":\"Title 60\"}}", resp.Body.String())
	// Test second update of post and incrementing id
	resp = api.Put("/api/v2/posts/60", map[string]string{
		"title":   "New Title",
		"content": "Lorem ipsum 2",
	})
	assert.Equal(t, http.StatusOK, resp.Code)
	require.JSONEq(t, "{\"result\":{\"author\":\"Homer\", \"content\":\"Lorem ipsum 2\", \"id\":60, \"title\":\"New Title\"}}", resp.Body.String())
	// Test with malformed author
	resp = api.Put("/api/v2/posts/70", map[string]string{
		"author": "New post ??",
	})
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	// Test with malformed id
	resp = api.Put("/api/v2/posts/aa", map[string]string{
		"content": "Lorem ipsum",
	})
	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
	// Test with large id
	resp = api.Put("/api/v2/posts/170", map[string]string{
		"author": "John Doe",
		"title":  "New post 1",
	})
	assert.Equal(t, http.StatusNotFound, resp.Code)
}
