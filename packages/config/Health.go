package config

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"Status"`
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	health := Health{Status: "healthy"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
