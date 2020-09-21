package Controller

import (
	"encoding/json"
	"net/http"
)

func respond(w http.ResponseWriter, status int, contentType string, payload interface{}) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
