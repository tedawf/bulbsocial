package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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
