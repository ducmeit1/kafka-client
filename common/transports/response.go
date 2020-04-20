package transports

import (
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, message map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(message)
	w.Write(b)
}

func BadRequest(w http.ResponseWriter, message map[string]interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(message)
	w.Write(b)
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(map[string]interface{}{
		"Error": "Internal Server Error",
	})
	w.Write(b)
}
