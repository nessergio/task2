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

func TestReadPost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupAppRouter(&TestConfig)
	statusCode, bodyString := unit.FakeRequest(t, "GET", "/api/v1/posts/15", router, nil)
	assert.Equal(t, http.StatusOK, statusCode)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      15,
		Title:   "Title 15",
		Content: "Neque velit amet labore adipisci. Etincidunt magnam etincidunt dolor velit numquam tempora. Amet quaerat amet tempora aliquam porro consectetur aliquam. Sit labore numquam etincidunt. Voluptatem modi quaerat dolorem.",
		Author:  "Author 15",
	})), bodyString)
	statusCode, _ = unit.FakeRequest(t, "GET", "/api/v1/posts/151", router, nil)
	assert.Equal(t, http.StatusNotFound, statusCode)
	statusCode, _ = unit.FakeRequest(t, "GET", "/api/v1/posts/asd", router, nil)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}
