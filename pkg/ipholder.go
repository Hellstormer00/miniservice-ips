package pkg

import (
	"log"
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
	IpChan      chan string
	VisitorChan chan int
	ips         hashset
}

func NewIpHolder(bufsize int) IpHolder {
	return IpHolder{
		IpChan:      make(chan string),
		VisitorChan: make(chan int),
		ips:         make(hashset),
	}
}

func (holder *IpHolder) Serve() {
	for {
		select {
		case ip := <-holder.IpChan:
			holder.addIp(ip)
		case holder.VisitorChan <- holder.getVisitors():
		}
	}
}

func (holder *IpHolder) addIp(new_ip string) {
	if added := holder.ips.add(new_ip); added {
		log.Printf("Added ip %s to ipHolder", new_ip)
	}
}

func (holder *IpHolder) getVisitors() int {
	return len(holder.ips)
}
