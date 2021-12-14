package pkg

import (
	"log"
	"net/http"
)

func HandleLogRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("New request:", r.Method, r.URL)
}

func HandleVisitorRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("New request:", r.Method, r.URL)
}
