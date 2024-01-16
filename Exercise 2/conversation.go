package main

import (
	"fmt"
	"net"
	"runtime"
)

func receiver(convChan chan string){
	// Specify the address to listen on (including port)
	recAddress := ":20013" //Opprett streng med IP addresse (eller portnummer, begge funker)
	udpAddr, err := net.ResolveUDPAddr("udp", recAddress) //finner UDP addresses tilknyttet IP addressen/portnummeret



	if err != nil { //evt si ifra hvis noe gikk galt
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	recConn, err := net.ListenUDP("udp", udpAddr) //returnerer enn net.UDPConn variabel som kan sende og receive fra UDP addressen
	if err != nil { //blablabla feilhåndtering
		fmt.Println("Error listening on UDP:", err)
		return
	}


	defer recConn.Close() //defer sier at en funksjon, i dette tilfellet conn.Close() som dreper UDPConn-en vår,
	//skal utføres når funksjonen den er i, i dette tilfellet main, er ferdig

	fmt.Println("UDP Server listening on", recAddress)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)
	
	for {
		// Read data from the connection
		n, clientAddr, err := recConn.ReadFromUDP(buffer) 
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Process the received data
		data := buffer[:n]
		fmt.Printf("Received from %v: %s\n", clientAddr, data)
		convChan <- ""
	}
}

func sender(convChan chan string){
	sendAddress := "20013"
	serverAddr, err := net.ResolveUDPAddr("udp", sendAddress)

	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	sendConn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}

	defer sendConn.Close()

	// Message to send
	message := []byte("Hiya guys!")

	// Send the message
	_, err = sendConn.Write(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Message sent to", serverAddr)
	convChan <- 
}


func main() {
	runtime.GOMAXPROCS(2)

	convChan := make(chan string)

	go sender(convChan)
	go receiver(convChan)
}