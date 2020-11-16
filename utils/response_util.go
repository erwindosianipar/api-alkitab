package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}

func ResponseMessage(w http.ResponseWriter, code int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}

func ResponseData(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}
