package main

import ( 
    "fmt"
) 

type IPAddr [4]byte

func (ip IPAddr) String() string {
	a := ip
	return string(ip[0]) + "." + string(a[1]) + "." + string(a[2]) + "." + string(a[3])
}	

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
