package main

import (
	"net/http"

	"github.com/hellstormer00/miniservice-ips/pkg"
)

const (
	bufsize = 10
)

func main() {
	ch := make(chan string, 10)
	ipHolder := pkg.IpHolder{
		Ips: make([]string, 0, bufsize),
	}
	http.HandleFunc("/logs", pkg.HandleLogRequest(ch, &ipHolder))
	http.HandleFunc("/visitors", pkg.HandleVisitorRequest)
	http.ListenAndServe(":5000", nil)
}
