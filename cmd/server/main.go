package main

import (
	"net/http"
	"sync"

	"github.com/hellstormer00/miniservice-ips/pkg"
)

const (
	bufsize = 50000
)

type IpHolder struct {
	mu  sync.Mutex
	ips []string
}

func (holder *IpHolder) AddIp(new_ip string) {
	holder.mu.Lock()

	result := false
	for _, ip := range holder.ips {
		if new_ip == ip {
			result = true
			break
		}
	}

	if result {
		holder.ips = append(holder.ips)
	}

	holder.mu.Unlock()
}

func main() {
	ch := make(chan string, 10)
	ipHolder := IpHolder{
		ips: make([]string, bufsize),
	}
	http.HandleFunc("/logs", pkg.HandleLogRequest(ch))
	http.HandleFunc("/visitors", pkg.HandleVisitorRequest)
	http.ListenAndServe(":5000", nil)
}
