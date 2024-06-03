package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task2/models/posts"
)

type PostsController struct {
	dataSource posts.DataSource
}

// listHandler responds with the list of all posts as JSON.
func (pc *PostsController) listHandler(c *gin.Context) {
	c.JSON(http.StatusOK, pc.dataSource.ListPosts())
}

// createHandler adds a post from JSON received in the request body.
func (pc *PostsController) createHandler(c *gin.Context) {
	var newPost posts.Post
	if err := c.BindJSON(&newPost); err != nil {
		_ = c.Error(err)
		_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid data object"))
		return
	}
	if err := newPost.Validate(); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pc.dataSource.CreatePost(&newPost)
	c.JSON(http.StatusCreated, &newPost)
}

// updateByIdHandler modifies a post with JSON received in the request body.
func (pc *PostsController) updateByIdHandler(c *gin.Context) {
	// retrieve id of the post from params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("invalid id"))
		return
	}
	// Find post we want to update
	post, err := pc.dataSource.ReadPost(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	// Try to update fields that are set in the command object
	if err = c.BindJSON(&post); err != nil {
		_ = c.Error(err)
		_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid data object"))
		return
	}
	// check if everything Ok with new values
	if err = post.Validate(); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if post, err = pc.dataSource.UpdatePost(post); err != nil {
		// we already checked for not found error, so this may not happen
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, &post)
}

// deleteByIdHandler deletes post by ID value sent by the client
func (pc *PostsController) deleteByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("invalid id"))
		return
	}
	post, err := pc.dataSource.DeletePost(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, post)
}

// readByIdHandler locates the post whose ID value matches the id
// parameter sent by the client, then returns that post as a response.
func (pc *PostsController) readByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("invalid id"))
		return
	}
	a, err := pc.dataSource.ReadPost(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

// NewPostsController creates new controller adding handlers to specified router group
// with specified data source
func NewPostsController(group *gin.RouterGroup, dataSource posts.DataSource) *PostsController {
	pc := PostsController{dataSource}
	group.GET("", pc.listHandler)
	group.GET(":id", pc.readByIdHandler)
	group.POST("", pc.createHandler)
	group.PUT(":id", pc.updateByIdHandler)
	group.DELETE(":id", pc.deleteByIdHandler)
	return &pc
}
