package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{true}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}
