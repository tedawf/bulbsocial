package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tedawf/tradebulb/internal/store"
)

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unauthorizedBasicError(w, r, fmt.Errorf("missing auth header"))
				return
			}

			// parse it -> get the base64
			header := strings.Split(authHeader, " ")
			if len(header) != 2 || header[0] != "Basic" {
				app.unauthorizedBasicError(w, r, fmt.Errorf("invalid auth header"))
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(header[1])
			if err != nil {
				app.unauthorizedBasicError(w, r, err)
				return
			}

			// check the credentials
			user := app.config.auth.basic.user
			pass := app.config.auth.basic.pass

			creds := strings.SplitN(string(decoded), ":", 2)
			if creds[0] != user || creds[1] != pass {
				app.unauthorizedBasicError(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (app *application) TokenAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unauthorizedBasicError(w, r, fmt.Errorf("missing auth header"))
				return
			}

			header := strings.Split(authHeader, " ")
			if len(header) != 2 || header[0] != "Bearer" {
				app.unauthorizedBasicError(w, r, fmt.Errorf("invalid auth header"))
				return
			}

			token := header[1]

			jwtToken, err := app.authenticator.ValidateToken(token)
			if err != nil {
				app.unauthorizedError(w, r, err)
				return
			}

			claims, _ := jwtToken.Claims.(jwt.MapClaims)

			userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
			if err != nil {
				app.unauthorizedError(w, r, err)
				return
			}

			ctx := r.Context()

			user, err := app.store.Users.GetByID(ctx, userID)
			if err != nil {
				app.unauthorizedError(w, r, err)
				return
			}

			ctx = context.WithValue(ctx, userCtx, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (app *application) checkPostOwnership(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromCtx(r)
		post := getPostFromCtx(r)

		// check if its the users post
		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		// check role precedence
		allowed, err := app.checkRolePrecedence(r.Context(), user, requiredRole)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		if !allowed {
			app.forbiddenError(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkRolePrecedence(ctx context.Context, user *store.User, roleName string) (bool, error) {
	role, err := app.store.Roles.GetByName(ctx, roleName)
	if err != nil {
		return false, err
	}

	return user.Role.Level >= role.Level, nil
}
