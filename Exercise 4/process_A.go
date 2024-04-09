package main

import (
	//"Network-go/network/bcast"
	//"Network-go/network/localip"
	//"Network-go/network/peers"
	//"flag"
	"fmt"
	"time"
)



func main() {
	//Termination after a couple of seconds
	port := "20013"
	adr := net.ResolveUDPAddr("udp", port)
	listener := net.ListenUDP("udp", adr)
	defer listener.Close()

	buffer := make([]byte,1024)
	for{
		n := listener.ReadFromUDP(buffer)
		data <- buffer[:n]			
		select()
	}

	fmt.Printf("Terminating...\n")

}
