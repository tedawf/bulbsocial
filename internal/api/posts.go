package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tedawf/bulbsocial/internal/db"
)

func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string   `json:"title" validate:"required,max=100"`
		Content string   `json:"content" validate:"required,max=1000"`
		Tags    []string `json:"tags"`
	}

	if err := readJSON(w, r, &req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	params := db.CreatePostParams{
		UserID:  50, // get postID from auth
		Title:   req.Title,
		Content: req.Content,
		Tags:    req.Tags,
	}

	post, err := s.postService.CreatePost(r.Context(), params)
	if err != nil {
		s.internalServerError(w, r, err)
		return
	}

	if err := s.jsonResponse(w, http.StatusCreated, post); err != nil {
		s.internalServerError(w, r, err)
		return
	}
}

func (s *Server) handleGetPost(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
	if err != nil {
		s.badRequestError(w, r, err)
		return
	}

	post, err := s.postService.GetPostByID(r.Context(), postID)
	if err != nil {
		s.notFoundError(w, r, err)
		return
	}

	if err := s.jsonResponse(w, http.StatusOK, post); err != nil {
		s.internalServerError(w, r, err)
	}
}
