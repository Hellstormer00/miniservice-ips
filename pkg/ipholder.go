package pkg

import (
	"log"
	"sync"
)

type IpHolder struct {
	mu  sync.Mutex
	ips []string
}

func NewIpHolder(bufsize int) IpHolder {
	return IpHolder{
		ips: make([]string, 0, bufsize),
	}
}

func (holder *IpHolder) AddIp(new_ip string) {
	holder.mu.Lock()
	defer holder.mu.Unlock()

	result := true
	for _, ip := range holder.ips {
		if new_ip == ip {
			result = false
			break
		}
	}

	if result {
		holder.ips = append(holder.ips, new_ip)
		log.Printf("Added ip %s to ipHolder", new_ip)
	}
}

func (holder *IpHolder) GetVisitors() int {
	holder.mu.Lock()
	defer holder.mu.Unlock()

	return len(holder.ips)
}
