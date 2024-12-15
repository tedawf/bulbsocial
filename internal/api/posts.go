package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tedawf/bulbsocial/internal/auth"
	"github.com/tedawf/bulbsocial/internal/db"
	"github.com/tedawf/bulbsocial/internal/service"
)

func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string `json:"title" validate:"required,max=100"`
		Content string `json:"content" validate:"required,max=1000"`
	}

	if err := s.parseAndValidate(w, r, &req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	authPayload := ctx.Value(authorizationPayloadKey).(*auth.Payload)

	params := db.CreatePostParams{
		UserID:  authPayload.UserID,
		Title:   req.Title,
		Content: req.Content,
	}

	post, err := s.postService.CreatePost(r.Context(), params)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			s.forbiddenError(w, r, err)
			return
		}
		s.internalServerError(w, r, err)
		return
	}

	if err := s.respond(w, http.StatusCreated, "created post successfully", post); err != nil {
		s.internalServerError(w, r, err)
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

	ctx := r.Context()
	authPayload := ctx.Value(authorizationPayloadKey).(*auth.Payload)

	if post.UserID != authPayload.UserID {
		err := errors.New("post does not belong to the authenticated user")
		s.unauthorizedError(w, r, err)
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
