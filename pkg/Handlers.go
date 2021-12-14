package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

type logRequest struct {
	TimeStamp string `json:"timestamp"`
	Ip        string `json:"ip"`
	Url       string `json:"url"`
}

func HandleLogRequest(ipHolder *IpHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request:", r.Method, r.URL)

		if r.Method != http.MethodPost {
			http.Error(w, "This API only uses POST", 400)
		}

		if r.Body == nil {
			http.Error(w, "No request body provided", 400)
		}

		log.Println("Decoding request body")
		var req logRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		ipHolder.AddIp(req.Ip)
	}
}

func HandleVisitorRequest(ipHolder *IpHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request:", r.Method, r.URL)

		if r.Method != http.MethodGet {
			http.Error(w, "This API only uses GET", 400)
		}

		log.Printf("Sending Response: %d ips", ipHolder.GetVisitors())
	}
}
