package main

import (
	"net/http"

	"github.com/hellstormer00/miniservice-ips/pkg"
)

const (
	bufsize = 50000
)

func main() {
	ch := make(chan string, 10)
	ipHolder := pkg.IpHolder{
		Ips: make([]string, bufsize),
	}
	http.HandleFunc("/logs", pkg.HandleLogRequest(ch))
	http.HandleFunc("/visitors", pkg.HandleVisitorRequest)
	http.ListenAndServe(":5000", nil)
}
