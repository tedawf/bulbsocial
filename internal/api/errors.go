package api

import (
	"net/http"
)

func (s *Server) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusInternalServerError, "the server encountered a problem", nil)
}
func (s *Server) forbiddenError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusForbidden, "forbidden", nil)
}
func (s *Server) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusBadRequest, err.Error(), nil)
}
func (s *Server) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.respond(w, http.StatusNotFound, "not found", nil)
}
