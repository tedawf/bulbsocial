package api

import (
	"net/http"
)

func (s *Server) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.jsonError(w, http.StatusInternalServerError, "the server encountered a problem")
}
func (s *Server) forbiddenError(w http.ResponseWriter, r *http.Request) {
	s.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path)

	s.jsonError(w, http.StatusForbidden, "forbidden")
}
func (s *Server) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.jsonError(w, http.StatusBadRequest, err.Error())
}
func (s *Server) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	s.jsonError(w, http.StatusNotFound, "not found")
}
