package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

var error_ch chan string
var sel = ""
var received_ch chan string
var quit chan bool
var i = 0

func init() {
	error_ch = make(chan string)
	received_ch = make(chan string)
	quit = make(chan bool)
}

func receiver() {
	// Specify the address to listen on (including port)
	recAddress := ":20013"                                //Opprett streng med IP addresse (eller portnummer, begge funker)
	udpAddr, err := net.ResolveUDPAddr("udp", recAddress) //finner UDP addresses tilknyttet IP addressen/portnummeret

	if err != nil { //evt si ifra hvis noe gikk galt
		error_ch <- "Error resolving UDP address:"
		return
	}

	// Create a UDP connection
	recConn, err := net.ListenUDP("udp", udpAddr) //returnerer enn net.UDPConn variabel som kan sende og receive fra UDP addressen
	if err != nil {                               //blablabla feilhåndtering
		error_ch <- "Error listening on UDP:"
		return
	}

	//defer recConn.Close() //defer sier at en funksjon, i dette tilfellet conn.Close() som dreper UDPConn-en vår,
	//skal utføres når funksjonen den er i, i dette tilfellet receiver, er ferdig

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, clientAddr, err := recConn.ReadFromUDP(buffer)
		if err != nil {
			error_ch <- "Error reading from UDP:"
			continue
		}

		// Process the received data
		data := buffer[:n]
		received_ch <- fmt.Sprintf("Received from %v: %s\n", clientAddr, data)

	}
}

func sender() {
	sendAddress := ":20013"
	serverAddr, err := net.ResolveUDPAddr("udp", sendAddress)

	if err != nil {
		error_ch <- "Error resolving UDP address:"
		return
	}

	sendConn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		error_ch <- "Error listening on UDP:"
		return
	}

	//defer sendConn.Close()

	// Message to send
	for {
		message := []byte("Hiya guys!")

		// Send the message
		_, err = sendConn.Write(message)
		if err != nil {
			error_ch <- "Error sending message:"
			return
		}
		time.Sleep(500 * time.Millisecond)
		i++
		if i == 10 {
			quit <- true
		}
	}

}

func server() {
	for {
		select {
		case sel := <-received_ch:
			fmt.Println(sel)
		case sel := <-error_ch:
			fmt.Println(sel)
			quit <- true
			break
		case <-quit:
			break
		}

	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go sender()
	go receiver()
	go server()

	<-quit

	fmt.Println("Jippi")
}
