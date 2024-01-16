package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve the UDP address of the server
	port := ":20013"
	serverAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	// Message to send
	message := []byte("Hiya guys!")

	// Send the message
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Message sent to", serverAddr)
}
