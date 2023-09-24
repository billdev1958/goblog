package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePost(*Post) error
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStore{
		db: db,
	}, nil
}

// Create post
func (s *PostgreStore) CreatePost(post *Post) error {
	query := `INSERT INTO posts
	(title, body, created_at)
	VALUES ($1, $2, $3)`

	_, err := s.db.Query(
		query,
		post.Title,
		post.Body,
		post.Created_at,
	)
	if err != nil {
		return err
	}
	return nil
}
