package main

import (
	"fmt"
	"net"
)

func main() {
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")

	fmt.Println(address)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, address)
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddresss, err := listener.ReadFromUDP(buffer)

		if err != nil {
			panic(err)
		}

		fmt.Printf("server %v\n", remoteAddresss)
		fmt.Printf("now %s\n", string(buffer[:length]))
	}
}
