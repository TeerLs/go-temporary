package res

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}