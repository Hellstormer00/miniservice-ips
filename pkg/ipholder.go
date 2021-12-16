package pkg

import (
	"log"
	"sync"
)

type hashset map[string]struct{}

func (set hashset) add(key string) bool {
	if _, ok := set[key]; !ok {
		set[key] = struct{}{}
		return true
	}
	return false
}

type IpHolder struct {
	mu  sync.Mutex
	ips hashset
}

func NewIpHolder(bufsize int) IpHolder {
	return IpHolder{
		ips: make(hashset),
	}
}

func (holder *IpHolder) AddIp(new_ip string) {
	holder.mu.Lock()
	defer holder.mu.Unlock()

	if added := holder.ips.add(new_ip); added {
		log.Printf("Added ip %s to ipHolder", new_ip)
	}
}

func (holder *IpHolder) GetVisitors() int {
	holder.mu.Lock()
	defer holder.mu.Unlock()

	return len(holder.ips)
}
