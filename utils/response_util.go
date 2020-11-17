package utils

import (
	"encoding/json"
	"net/http"
)

// Message is commons to encode response message
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}

// ResponseMessage is commons to encode response message with data
func ResponseMessage(w http.ResponseWriter, code int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}

// ResponseData is commons to encode response data
func ResponseData(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}
