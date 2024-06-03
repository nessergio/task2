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

func TestReadPost(t *testing.T) {
	_, api := humatest.New(t)
	postsMemory := posts.FromFileToMemoryDs(TestConfig.InitialDataFile)
	v2.NewPostsController(api, "/api/v2/posts", postsMemory)
	gin.SetMode(gin.ReleaseMode)
	resp := api.Get("/api/v2/posts/15")
	assert.Equal(t, http.StatusOK, resp.Code)
	require.JSONEq(t, "{\"result\":{\"id\":15,\"title\":\"Title 15\",\"content\":\"Neque velit amet labore adipisci. Etincidunt magnam etincidunt dolor velit numquam tempora. Amet quaerat amet tempora aliquam porro consectetur aliquam. Sit labore numquam etincidunt. Voluptatem modi quaerat dolorem.\",\"author\":\"Author 15\"}}", resp.Body.String())
	resp = api.Get("/api/v2/posts/151")
	assert.Equal(t, http.StatusNotFound, resp.Code)
	resp = api.Get("/api/v2/posts/asd")
	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}
