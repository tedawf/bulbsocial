package api

import (
	"net/http"
)

func (s *Server) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusInternalServerError, "the server encountered a problem", err.Error())
}
func (s *Server) forbiddenError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusForbidden, "forbidden", err.Error())
}
func (s *Server) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusBadRequest, "bad request", err.Error())
}
func (s *Server) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusNotFound, "not found", err.Error())
}
func (s *Server) unauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusUnauthorized, "unauthorized", err.Error())
}
func (s *Server) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusConflict, "conflict", err.Error())
}
