package main

import (
	"fmt"
	"net"
)

func main() {
	// Specify the address to listen on (including port)
	address := ":30000"                                //":30000" //Opprett streng med IP addresse (eller portnummer, begge funker)
	udpAddr, err := net.ResolveUDPAddr("udp", address) //finner UDP addresses tilknyttet IP addressen/portnummeret

	if err != nil { //evt si ifra hvis noe gikk galt
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	secconn, err := net.ListenUDP("udp", udpAddr) //returnerer en net.UDPConn variabel som kan sende og receive fra UDP addressen

	if err != nil { //blablabla feilhåndtering
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer secconn.Close() //defer sier at en funksjon, i dette tilfellet conn.Close() som dreper UDPConn-en vår,
	//skal utføres når funksjonen den er i, i dette tilfellet main, er ferdig

	fmt.Println("UDP Server listening on", address)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	// Read data from the connection
	n, clientAddr, err := secconn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		return
	}

	// Process the received data
	data := buffer[:n]
	fmt.Printf("Received from %v: %s\n", clientAddr, data)

}
