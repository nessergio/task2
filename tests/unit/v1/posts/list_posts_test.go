package posts

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"task2/models/posts"
	"task2/routes"
	"task2/tests/unit"
	"testing"
)

func TestListPosts(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupAppRouter(&TestConfig)
	statusCode, bodyString := unit.FakeRequest(t, "GET", "/api/v1/posts", router, nil)
	assert.Equal(t, http.StatusOK, statusCode)
	var postsList []posts.Post
	err := json.Unmarshal([]byte(bodyString), &postsList)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(postsList), 100)
	require.JSONEq(t, string(unit.AssertMarshallJson(t, posts.Post{
		ID:      5,
		Title:   "Title 5",
		Content: "Consectetur labore dolorem ut dolore amet. Ut quaerat ut porro quisquam dolor. Neque modi ipsum dolor ut tempora quaerat sed. Tempora adipisci ut eius ut. Tempora tempora etincidunt sed tempora magnam.",
		Author:  "Author 5",
	})), string(unit.AssertMarshallJson(t, postsList[4])))
}
