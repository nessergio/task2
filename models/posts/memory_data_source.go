package posts

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
)

type MemoryDataSource struct {
	Posts     map[int]*Post
	MaxPostId int
}

func errorNotFound(id int) error {
	return fmt.Errorf("post with id %d is not found", id)
}

// FromFileToMemoryDs reads mock data from JSON file
func FromFileToMemoryDs(fileName string) *MemoryDataSource {
	pd := MemoryDataSource{make(map[int]*Post), 0}
	if fileName == "" {
		// if no file specified, return empty map
		return &pd
	}

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Printf("can not open data file %s: %v", fileName, err)
	} else {
		log.Printf("successfully read blog data file '%s'", fileName)
		// defer the closing of our jsonFile so that we can parse it later on
		defer func() {
			if err = jsonFile.Close(); err != nil {
				log.Fatalf("can not close file: %v", err)
			}
		}()
	}

	var data struct {
		Posts []Post `json:"posts"`
	}

	err = json.NewDecoder(jsonFile).Decode(&data)
	if err != nil {
		log.Printf("error reading data file, initial dataset will be empty: %v", err)
	}

	// moving posts from slice to map for better random access
	for i, p := range data.Posts {
		pd.Posts[p.ID] = &data.Posts[i]
		if p.ID > pd.MaxPostId {
			pd.MaxPostId = p.ID
		}
	}

	return &pd
}

// ListPosts converts map of post to the list
func (pm *MemoryDataSource) ListPosts() []*Post {
	res := make([]*Post, 0, len(pm.Posts))
	for _, p := range pm.Posts {
		res = append(res, p)
	}
	// Sort by post ID
	slices.SortFunc(res, func(a, b *Post) int {
		return cmp.Compare(a.ID, b.ID)
	})
	return res
}

// CreatePost adds post to the map and sets id with maximum index
func (pm *MemoryDataSource) CreatePost(post *Post) (*Post, error) {
	pm.MaxPostId++
	post.ID = pm.MaxPostId
	pm.Posts[post.ID] = post
	return post, nil
}

// ReadPost retrieves post  with given id from the map
func (pm *MemoryDataSource) ReadPost(id int) (*Post, error) {
	post, ok := pm.Posts[id]
	if !ok {
		return nil, errorNotFound(id)
	}
	return post, nil
}

// UpdatePost replaces record with given id in the map
func (pm *MemoryDataSource) UpdatePost(post *Post) (*Post, error) {
	_, ok := pm.Posts[post.ID]
	if !ok {
		return nil, errorNotFound(post.ID)
	}
	pm.Posts[post.ID] = post
	return post, nil
}

// DeletePost deletes record with given id form a map
func (pm *MemoryDataSource) DeletePost(id int) (*Post, error) {
	post, ok := pm.Posts[id]
	if !ok {
		return nil, errorNotFound(id)
	}
	delete(pm.Posts, id)
	return post, nil
}
