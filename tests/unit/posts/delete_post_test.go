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

func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupAppRouter(&TestConfig)
	statusCode, bodyString := unit.FakeRequest(t, "DELETE", "/api/v1/posts/25", router, nil)
	assert.Equal(t, http.StatusOK, statusCode)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      25,
		Title:   "Title 25",
		Content: "Sed ipsum aliquam velit est. Modi est modi quisquam. Amet numquam numquam adipisci quaerat consectetur. Porro voluptatem consectetur dolorem dolore ipsum voluptatem. Labore porro magnam etincidunt ut sed sit etincidunt.",
		Author:  "Author 25",
	})), bodyString)
	statusCode, _ = unit.FakeRequest(t, "DELETE", "/api/v1/posts/25", router, nil)
	assert.Equal(t, http.StatusNotFound, statusCode)
	statusCode, _ = unit.FakeRequest(t, "DELETE", "/api/v1/posts/250", router, nil)
	assert.Equal(t, http.StatusNotFound, statusCode)
	statusCode, _ = unit.FakeRequest(t, "DELETE", "/api/v1/posts/asd", router, nil)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}
