package pkg

import (
	"log"
	"sync"
)

type IpHolder struct {
	mu  sync.Mutex
	Ips []string
}

func (holder *IpHolder) AddIp(new_ip string) {
	holder.mu.Lock()

	result := true
	for _, ip := range holder.Ips {
		if new_ip == ip {
			result = false
			break
		}
	}

	if result {
		holder.Ips = append(holder.Ips, new_ip)
		log.Printf("Added ip %s to ipHolder", new_ip)
	}

	holder.mu.Unlock()
}

func (holder *IpHolder) GetVisitors() int {
	return len(holder.Ips)
}
