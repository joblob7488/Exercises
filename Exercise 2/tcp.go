package main

import (
	"fmt"
	"net"
)

type msg struct {
	text string
}

func transmit(conn net.Conn, err error, tx_data chan []byte) {
	//dial up the server, establish connection:

	fmt.Println("Status: ", err)

	//send message
	if err == nil {
		conn.Write(<-tx_data)
		fmt.Println("Sent")
	} else {
		fmt.Println("Error: ", err)
	}

}

func receive(conn net.Conn, rx_data chan []byte) {
	//read from buffer
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		rx_data <- buffer[:n]
	}

	defer conn.Close()
}

func main() {
	var adress string = "10.100.23.129:33546"
	var network_type string = "tcp"
	data := []byte("Hello mister")

	rx_data := make(chan []byte)
	tx_data := make(chan []byte)
	conn, err := net.Dial(network_type, adress)

	go transmit(conn, err, tx_data)
	go receive(conn, rx_data)

	tx_data <- data

	for {
		select {
		case in := <-rx_data:
			fmt.Println("Recieved: ", in)
		}
	}

}
