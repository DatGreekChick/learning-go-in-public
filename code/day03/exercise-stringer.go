package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var ipString string
	for i := range ip {
		if i == len(ip)-1 {
			ipString += fmt.Sprint(ip[i])
		} else {
			ipString += fmt.Sprint(ip[i]) + "."
		}
	}
	return ipString
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
		// will print the following for each key:
		// loopback: 127.0.0.1
		// googleDNS: 8.8.8.8
	}
}
