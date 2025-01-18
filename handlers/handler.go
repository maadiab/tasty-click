package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/maadiab/tasty-click/database"
)

type ApiConfig struct {
	DBQuries database.Queries
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, statusCode /*code*/ int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Code:    statusCode,
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write error!!!, ", http.StatusInternalServerError)
	}
}
