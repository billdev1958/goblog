package db

import "time"

type User struct {
	ID    int
	Name  string
	Email string
}

type Post struct {
	ID         int
	Title      string
	Body       string
	Created_at time.Time
}

type CreatePostRequest struct {
	Title string
	Body  string
}

func NewPost(title, body string) (*Post, error) {
	return &Post{
		Title:      title,
		Body:       body,
		Created_at: time.Now().UTC(),
	}, nil
}
