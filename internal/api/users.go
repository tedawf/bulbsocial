package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
	"github.com/tedawf/bulbsocial/internal/auth"
	"github.com/tedawf/bulbsocial/internal/db"
	"github.com/tedawf/bulbsocial/internal/service"
)

type userResponse struct {
	ID                int64        `json:"id"`
	Email             string       `json:"email"`
	Username          string       `json:"username"`
	CreatedAt         time.Time    `json:"created_at"`
	PasswordChangedAt sql.NullTime `json:"password_changed_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
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

	res := newUserResponse(user)
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

	if err := s.parseAndValidate(w, r, &req); err != nil {
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

	res := newUserResponse(user)
	if err := s.respond(w, http.StatusCreated, "created user successfully", res); err != nil {
		s.internalServerError(w, r, err)
	}
}

func (s *Server) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username" validate:"required,alphanum"`
		Password string `json:"password" validate:"required,min=6"`
	}

	if err := s.parseAndValidate(w, r, &req); err != nil {
		s.badRequestError(w, r, err)
		return
	}

	user, accessToken, err := s.userService.LoginUser(r.Context(), req.Username, req.Password, s.config.AccessTokenDuration)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			s.unauthorizedError(w, r, err)
			return
		}
		s.internalServerError(w, r, err)
		return
	}

	res := struct {
		AccessToken string       `json:"access_token"`
		User        userResponse `json:"user"`
	}{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	if err := s.respond(w, http.StatusOK, "user login successfully", res); err != nil {
		s.internalServerError(w, r, err)
	}
}

func (s *Server) handleFollowUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		s.badRequestError(w, r, fmt.Errorf("invalid user ID: %w", err))
		return
	}

	ctx := r.Context()
	authPayload := ctx.Value(authorizationPayloadKey).(*auth.Payload)

	if authPayload.UserID == userID {
		err = errors.New("cannot follow yourself")
		s.badRequestError(w, r, err)
		return
	}

	err = s.userService.FollowUser(ctx, authPayload.UserID, userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			s.notFoundError(w, r, err)
			return
		}
		if errors.Is(err, service.ErrAlreadyFollowing) {
			s.conflictError(w, r, err)
			return
		}
		s.internalServerError(w, r, err)
		return
	}

	if err := s.respond(w, http.StatusCreated, "success", "followed user successfully"); err != nil {
		s.internalServerError(w, r, err)
	}
}

func (s *Server) handleUnfollowUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		s.badRequestError(w, r, fmt.Errorf("invalid user ID: %w", err))
		return
	}

	ctx := r.Context()
	authPayload := ctx.Value(authorizationPayloadKey).(*auth.Payload)

	err = s.userService.UnfollowUser(ctx, authPayload.UserID, userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			s.notFoundError(w, r, err)
			return
		}
		if errors.Is(err, service.ErrNotFollowing) {
			s.conflictError(w, r, err)
			return
		}
		s.internalServerError(w, r, err)
		return
	}

	if err := s.respond(w, http.StatusOK, "success", "unfollowed user successfully"); err != nil {
		s.internalServerError(w, r, err)
	}
}
