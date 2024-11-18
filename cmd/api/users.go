package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tedawf/tradebulb/internal/store"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil || userID < 1 {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.fetchUser(ctx, userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

type Follow struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUser(w http.ResponseWriter, r *http.Request) {
	followedUserID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user := getUserFromCtx(r)

	if err := app.store.Followers.Follow(r.Context(), user.ID, followedUserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictError(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) unfollowUser(w http.ResponseWriter, r *http.Request) {
	unfollowedUserID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user := getUserFromCtx(r)

	if err := app.store.Followers.Unfollow(r.Context(), user.ID, unfollowedUserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) verifyUser(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	err := app.store.Users.Verify(r.Context(), token)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
