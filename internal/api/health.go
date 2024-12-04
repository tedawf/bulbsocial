package api

import "net/http"

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := s.respond(w, http.StatusOK, "ok", nil); err != nil {
		s.internalServerError(w, r, err)
	}
}
