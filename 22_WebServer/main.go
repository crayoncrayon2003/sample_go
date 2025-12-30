package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type StatusResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
	Uptime string `json:"uptime"`
}

var startTime time.Time

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := StatusResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
		Uptime: time.Since(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func main() {
	startTime = time.Now()

	http.HandleFunc("/status", statusHandler)

	fmt.Println("Web server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
