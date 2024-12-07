package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
	"github.com/tedawf/bulbsocial/internal/db"
)

type UserResponse struct {
	ID                int64        `json:"id"`
	Email             string       `json:"email"`
	Username          string       `json:"username"`
	CreatedAt         time.Time    `json:"created_at"`
	PasswordChangedAt sql.NullTime `json:"password_changed_at"`
}

func NewUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Email:             user.Email,
		Username:          user.Username,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil || userID < 1 {
		s.badRequestError(w, r, fmt.Errorf("invalid user ID: must be a positive integer"))
		return
	}

	user, err := s.userService.GetUserByID(r.Context(), userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			s.notFoundError(w, r, err)
		default:
			s.internalServerError(w, r, err)
		}
		return
	}

	res := NewUserResponse(user)

	if err := s.respond(w, http.StatusOK, "fetched user successfully", res); err != nil {
		s.internalServerError(w, r, err)
	}
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username" validate:"required,alphanum"` // todo: get postID from auth
		Password string `json:"password" validate:"required,min=6"`
		Email    string `json:"email" validate:"required,email"`
	}

	if err := s.parse(w, r, &req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	user, err := s.userService.CreateUser(r.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				s.forbiddenError(w, r, err)
				return
			}
		}
		s.internalServerError(w, r, err)
		return
	}

	res := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	if err := s.respond(w, http.StatusCreated, "created user successfully", res); err != nil {
		s.internalServerError(w, r, err)
		return
	}
}
