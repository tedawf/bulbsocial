package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type APIResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func (s *Server) respond(w http.ResponseWriter, status int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := APIResponse[interface{}]{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return json.NewEncoder(w).Encode(response)
}

func (s *Server) parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func (s *Server) parseAndValidate(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := s.parse(w, r, data); err != nil {
		return fmt.Errorf("error parsing request body: %w", err)
	}

	if err := Validate.Struct(data); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
