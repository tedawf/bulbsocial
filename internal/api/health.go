package api

import "net/http"

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := s.jsonResponse(w, http.StatusOK, "ok"); err != nil {
		s.internalServerError(w, r, err)
	}
}
