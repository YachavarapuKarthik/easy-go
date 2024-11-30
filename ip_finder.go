package main 

import (
	"fmt"
	"net"
)

func ip_finder(){
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		ip, ok := addr.(*net.IPNet)
		if ok && ip.IP.To4() != nil {
			fmt.Println(ip.IP.String())
		}
	}
}