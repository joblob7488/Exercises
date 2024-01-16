package main

import (
	"fmt"
	"net"
)

func main() {
	// Specify the address to listen on (including port)
	address := ":20013"                                //":30000" //Opprett streng med IP addresse (eller portnummer, begge funker)
	udpAddr, err := net.ResolveUDPAddr("udp", address) //finner UDP addresses tilknyttet IP addressen/portnummeret

	if err != nil { //evt si ifra hvis noe gikk galt
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr) //returnerer en net.UDPConn variabel som kan sende og receive fra UDP addressen

	if err != nil { //blablabla feilhåndtering
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP Server dialing on", address)

	// message er strengen som sendes til serveren
	message := []byte("Hello Mario")

	_, err = conn.Write(message) //sender message til serveren
	if err != nil {
		fmt.Println("Error writing to UDP address:", err)
		return
	}

	fmt.Println("Data sent to the UDP server.")
}
