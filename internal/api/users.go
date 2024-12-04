package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		s.badRequestError(w, r, err)
		return
	}

	user, err := s.userService.GetUserByID(r.Context(), userID)
	if err != nil {
		s.notFoundError(w, r, err)
		return
	}

	type UserResponse struct {
		ID         int64     `json:"id"`
		Email      string    `json:"email"`
		Username   string    `json:"username"`
		CreatedAt  time.Time `json:"created_at"`
		IsVerified bool      `json:"is_verified"`
		RoleID     int32     `json:"role_id"`
	}

	res := &UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		IsVerified: user.IsVerified,
		RoleID:     user.RoleID,
	}

	if err := s.jsonResponse(w, http.StatusOK, res); err != nil {
		s.internalServerError(w, r, err)
	}
}
