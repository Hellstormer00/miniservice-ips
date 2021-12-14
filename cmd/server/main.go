package main

import (
	"net/http"

	"github.com/hellstormer00/miniservice-ips/pkg"
)

func main() {
	http.HandleFunc("/logs", pkg.HandleLogRequest)
	http.HandleFunc("/visitors", pkg.HandleVisitorRequest)
	http.ListenAndServe(":5000", nil)
}
