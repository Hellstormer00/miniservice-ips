package main

import (
	"net/http"

	"github.com/hellstormer00/miniservice-ips/pkg"
)

const (
	bufsize = 10
)

func main() {
	ipHolder := pkg.IpHolder{
		Ips: make([]string, 0, bufsize),
	}
	http.HandleFunc("/logs", pkg.HandleLogRequest(&ipHolder))
	http.HandleFunc("/visitors", pkg.HandleVisitorRequest(&ipHolder))
	http.ListenAndServe(":5000", nil)
}
