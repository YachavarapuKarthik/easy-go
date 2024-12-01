package main

import (
	"fmt"
	"net"
)

func dns_to_ip() {
	// Define the domain name you want to resolve
	domain := "localhost"

	// Use net.LookupIP to get the IP addresses of the domain
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("Error resolving domain: %v\n", err)
		return
	}

	// Print the IP addresses
	fmt.Printf("IP addresses for domain %s:\n", domain)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
