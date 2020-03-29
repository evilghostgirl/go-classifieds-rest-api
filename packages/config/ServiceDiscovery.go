package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	BASE_CONSUL_URL = "http://consulz:8500"
	PORT            = 8080
)

type RegisterRequest struct {
	Name    string    `json:"Name"`
	Address string    `json:"Address"`
	Check   CheckInfo `json:"Check"`
}

type CheckInfo struct {
	Http     string `json:"Http"`
	Interval string `json:"Interval"`
}

func RegisterInServiceDiscovery() {
	url := BASE_CONSUL_URL + "/v1/agent/service/register"
	objectdata := RegisterRequest{
		Name:    "GolangApp",
		Address: "app",
		Check: CheckInfo{
			Http:     "http://app:3000/health",
			Interval: "10s",
		},
	}

	client := &http.Client{}

	body, err := json.Marshal(objectdata)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	fmt.Print("registered")

	fmt.Println(resp.StatusCode)

}
