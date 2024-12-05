package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
	"github.com/tedawf/bulbsocial/internal/db"
)

func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID  int64  `json:"user_id" validate:"gt=0"` // todo: get postID from auth
		Title   string `json:"title" validate:"required,max=100"`
		Content string `json:"content" validate:"required,max=1000"`
	}

	if err := s.parse(w, r, &req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	params := db.CreatePostParams{
		UserID:  req.UserID,
		Title:   req.Title,
		Content: req.Content,
	}

	post, err := s.postService.CreatePost(r.Context(), params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				s.forbiddenError(w, r, err)
				return
			}
		}
		s.internalServerError(w, r, err)
		return
	}

	if err := s.respond(w, http.StatusCreated, "created post successfully", post); err != nil {
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

	if err := s.respond(w, http.StatusOK, "fetched post successfully", post); err != nil {
		s.internalServerError(w, r, err)
	}
}

func (s *Server) handleGetFeed(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	posts, err := s.postService.GetAllPosts(r.Context(), int32(limit), int32(offset))
	if err != nil {
		s.notFoundError(w, r, err)
		return
	}

	if err := s.respond(w, http.StatusOK, "fetched all posts successfully", posts); err != nil {
		s.internalServerError(w, r, err)
	}
}
