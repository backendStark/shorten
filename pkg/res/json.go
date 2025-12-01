package res

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
