package main

import (
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

// Ping make a ping to an address and retun true of false
func Ping(ipAddress string) (bool, error) {
	ret := false

	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ipAddress)
	if err != nil {
		return false, err
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		ret = true
	}
	// p.OnIdle = func() {
	// 	fmt.Println("finish")
	// }
	err = p.Run()
	if err != nil {
		return false, err
	}
	return ret, nil
}
