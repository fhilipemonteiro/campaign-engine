package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, code string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := map[string]interface{}{
		"code": code,
		"data": data,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to write JSON response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
