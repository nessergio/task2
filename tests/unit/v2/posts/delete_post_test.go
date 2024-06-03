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

func TestDeletePost(t *testing.T) {
	_, api := humatest.New(t)
	postsMemory := posts.FromFileToMemoryDs(TestConfig.InitialDataFile)
	v2.NewPostsController(api, "/api/v2/posts", postsMemory)
	gin.SetMode(gin.ReleaseMode)
	resp := api.Delete("/api/v2/posts/25")
	assert.Equal(t, http.StatusOK, resp.Code)
	require.JSONEq(t, "{\"result\":{\"author\":\"Author 25\", \"content\":\"Sed ipsum aliquam velit est. Modi est modi quisquam. Amet numquam numquam adipisci quaerat consectetur. Porro voluptatem consectetur dolorem dolore ipsum voluptatem. Labore porro magnam etincidunt ut sed sit etincidunt.\", \"id\":25, \"title\":\"Title 25\"}}",
		resp.Body.String())
	resp = api.Delete("/api/v2/posts/25")
	assert.Equal(t, http.StatusNotFound, resp.Code)
	resp = api.Delete("/api/v2/posts/250")
	assert.Equal(t, http.StatusNotFound, resp.Code)
	resp = api.Delete("/api/v2/posts/asd")
	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}
