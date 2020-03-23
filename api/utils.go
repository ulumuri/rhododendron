package api

import (
	"encoding/json"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
