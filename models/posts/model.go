package posts

import (
	"fmt"
	"strings"
)

type Post struct {
	ID      int    `json:"id" minimum:"0"`
	Title   string `json:"title" maxLength:"50"`
	Content string `json:"content" maxLength:"5000"`
	Author  string `json:"author" maxLength:"50"`
}

type DataSource interface {
	ListPosts() []*Post
	CreatePost(*Post) (*Post, error)
	ReadPost(int) (*Post, error)
	UpdatePost(*Post) (*Post, error)
	DeletePost(int) (*Post, error)
}

const prohibitedSymbols = "!?@#$%^&'\"\\\r\n"

func (p *Post) Validate() error {
	if len(p.Title) == 0 {
		return fmt.Errorf("title can`t be empty")
	}
	if len(p.Content) == 0 {
		return fmt.Errorf("content can`t be empty")
	}
	if len(p.Author) == 0 {
		return fmt.Errorf("author can`t be empty")
	}
	if len(p.Title) > 50 {
		return fmt.Errorf("title is too long: %d. Must be < 50", len(p.Title))
	}
	if len(p.Author) > 50 {
		return fmt.Errorf("author name is too long: %d. Must be < 50", len(p.Author))
	}
	if len(p.Content) > 5000 {
		return fmt.Errorf("content is too long: %d. Must be < 5000", len(p.Content))
	}
	if strings.ContainsAny(p.Author, prohibitedSymbols) {
		return fmt.Errorf("author name contains special characters")
	}
	if strings.ContainsAny(p.Title, prohibitedSymbols) {
		return fmt.Errorf("title contains special characters")
	}
	if strings.ContainsAny(p.Content, prohibitedSymbols) {
		return fmt.Errorf("content contains special characters")
	}
	return nil
}
