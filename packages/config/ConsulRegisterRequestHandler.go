package config

import (
	"encoding/json"
	"net/http"
)

// const (
// 	BASE_CONSUL_URL = "http://consulz:8500"
// 	PORT            = 8080
// )

// type RegisterRequest struct {
// 	Name    string    `json:"Name"`
// 	Address string    `json:"Address"`
// 	Check   CheckInfo `json:"Check"`
// }

// type CheckInfo struct {
// 	Http     string `json:"Http"`
// 	Interval string `json:"Interval"`
// }

func GetConsulRegisterRequest(w http.ResponseWriter, r *http.Request) {
	objectdata := RegisterRequest{
		Name:    "GolangApp",
		Address: "app",
		Check: CheckInfo{
			Http:     "http://app:3000/health",
			Interval: "10s",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(objectdata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
