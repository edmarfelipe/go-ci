package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// ErrorResponse is a generic error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteBadRequest writes a bad request response.
func WriteBadRequest(w http.ResponseWriter, err string) {
	WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err})
}

// WriteNotFound writes a not found response.
func WriteNotFound(w http.ResponseWriter, err string) {
	WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: err})
}

// WriteUnauthorized writes an unauthorized response.
func WriteUnauthorized(w http.ResponseWriter, err string) {
	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: err})
}

// WriteInternalError logs the error and writes a generic error response.
func WriteInternalError(w http.ResponseWriter, err error) {
	slog.Error("internal error", "err", err.Error())
	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
}

// WriteJSON writes a JSON response.
func WriteJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

// DecodeJSON decodes a JSON request body.
func DecodeJSON[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}
