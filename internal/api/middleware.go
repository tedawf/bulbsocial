package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (s *Server) tokenAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get(authorizationHeaderKey)
			if authHeader == "" {
				s.unauthorizedError(w, r, fmt.Errorf("missing auth header"))
				return
			}

			fields := strings.Fields(authHeader)
			if len(fields) != 2 || strings.ToLower(fields[0]) != authorizationTypeBearer {
				s.unauthorizedError(w, r, fmt.Errorf("invalid auth header"))
				return
			}

			accessToken := fields[1]

			payload, err := s.tokenMaker.VerifyToken(accessToken)
			if err != nil {
				s.unauthorizedError(w, r, err)
				return
			}

			ctx := r.Context()

			ctx = context.WithValue(ctx, authorizationPayloadKey, payload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
