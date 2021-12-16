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
			http.Error(w, "This API only uses POST", http.StatusMethodNotAllowed)
			return
		}

		if r.Body == nil {
			http.Error(w, "No request body provided", http.StatusBadRequest)
			return
		}

		log.Println("Decoding request body")
		var req logRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ipHolder.IpChan <- req.Ip
	}
}

func HandleVisitorRequest(ipHolder *IpHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request:", r.Method, r.URL)

		if r.Method != http.MethodGet {
			http.Error(w, "This API only uses GET", http.StatusMethodNotAllowed)
			return
		}

		log.Printf("Sending Response: %d ips", <-ipHolder.VisitorChan)

		b, err := json.Marshal(<-ipHolder.VisitorChan)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(b)
	}
}
