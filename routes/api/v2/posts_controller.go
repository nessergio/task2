package v2

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"task2/models/posts"
)

type PostsController struct {
	dataSource posts.DataSource
}

type PostsList struct {
	Body struct {
		Result []*posts.Post `json:"result" doc:"List of posts"`
	}
}

// listHandler responds with the list of all posts as JSON.
func (pc *PostsController) listHandler(ctx context.Context, input *struct{}) (*PostsList, error) {
	o := PostsList{}
	o.Body.Result = pc.dataSource.ListPosts()
	return &o, nil
}

type PostSingle struct {
	Body struct {
		Result *posts.Post `json:"result" doc:"Single post"`
	}
}

type IdInput struct {
	Id int `path:"id" minimum:"0" example:"100" doc:"ID of the post"`
}

// readByIdHandler locates the post whose ID value matches the id
// parameter sent by the client, then returns that post as a response.
func (pc *PostsController) readByIdHandler(ctx context.Context, input *IdInput) (*PostSingle, error) {
	res := PostSingle{}
	post, err := pc.dataSource.ReadPost(input.Id)
	if err != nil {
		return nil, huma.Error404NotFound("post not found", err)
	}
	res.Body.Result = post
	return &res, nil
}

// deleteByIdHandler deletes post by ID value sent by the client
func (pc *PostsController) deleteByIdHandler(ctx context.Context, input *IdInput) (*PostSingle, error) {
	res := PostSingle{}
	post, err := pc.dataSource.DeletePost(input.Id)
	if err != nil {
		return nil, huma.Error404NotFound("post not found", err)
	}
	res.Body.Result = post
	return &res, nil
}

// CreatePostCommand represents the create post request.
type CreatePostCommand struct {
	Body struct {
		Author  string `json:"author" maxLength:"50" doc:"Author of the post"`
		Title   string `json:"title" maxLength:"50" doc:"Title of the post"`
		Content string `json:"content" maxLength:"5000" doc:"Content"`
	}
}

// createHandler adds a post from JSON received in the request body.
func (pc *PostsController) createHandler(ctx context.Context, cmd *CreatePostCommand) (*PostSingle, error) {
	newPost := posts.Post{
		Author:  cmd.Body.Author,
		Title:   cmd.Body.Title,
		Content: cmd.Body.Content,
	}
	if err := newPost.Validate(); err != nil {
		return nil, huma.Error400BadRequest(err.Error(), err)
	}
	res := PostSingle{}
	post, err := pc.dataSource.CreatePost(&newPost)
	res.Body.Result = post
	return &res, err
}

// UpdatePostCommand represents the update post request.
type UpdatePostCommand struct {
	Id   int `path:"id" minimum:"0" example:"100" doc:"ID of the post"`
	Body struct {
		Author  string `json:"author,omitempty" maxLength:"50" doc:"Author of the post"`
		Title   string `json:"title,omitempty" maxLength:"50" doc:"Title of the post"`
		Content string `json:"content,omitempty" maxLength:"5000" doc:"Content"`
	}
}

// updateByIdHandler modifies a post with JSON received in the request body.
func (pc *PostsController) updateByIdHandler(ctx context.Context, input *UpdatePostCommand) (*PostSingle, error) {
	// Find post we want to update
	post, err := pc.dataSource.ReadPost(input.Id)
	if err != nil {
		return nil, err
	}
	if input.Body.Title != "" {
		post.Title = input.Body.Title
	}
	if input.Body.Author != "" {
		post.Author = input.Body.Author
	}
	if input.Body.Content != "" {
		post.Content = input.Body.Content
	}
	if err = post.Validate(); err != nil {
		return nil, huma.Error400BadRequest(err.Error(), err)
	}
	res := PostSingle{}
	post, err = pc.dataSource.UpdatePost(post)
	res.Body.Result = post
	return &res, err
}

// NewPostsController creates new controller adding handlers to specified router group
// with specified data source
func NewPostsController(api huma.API, baseUrl string, dataSource posts.DataSource) *PostsController {
	pc := PostsController{dataSource}
	huma.Get(api, baseUrl, pc.listHandler)
	huma.Get(api, baseUrl+"/{id}", pc.readByIdHandler)
	huma.Post(api, baseUrl, pc.createHandler)
	huma.Put(api, baseUrl+"/{id}", pc.updateByIdHandler)
	huma.Delete(api, baseUrl+"/{id}", pc.deleteByIdHandler)
	return &pc
}
