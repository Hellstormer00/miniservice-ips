package pkg

import "sync"

type IpHolder struct {
	mu  sync.Mutex
	Ips []string
}

func (holder *IpHolder) AddIp(new_ip string) {
	holder.mu.Lock()

	result := false
	for _, ip := range holder.Ips {
		if new_ip == ip {
			result = true
			break
		}
	}

	if result {
		holder.Ips = append(holder.Ips)
	}

	holder.mu.Unlock()
}
