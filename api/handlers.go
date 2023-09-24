package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billdev1958/goblog.git/db"
)

// handler Main
func (s *Server) handlePosts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreatePost(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handler post article
func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	req := new(db.CreatePostRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	post, err := db.NewPost(
		req.Title,
		req.Body)
	if err != nil {
		return err
	}
	if err := s.store.CreatePost(post); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, post)
}

// structure handlerfunc
type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

// Make handler function for api
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
