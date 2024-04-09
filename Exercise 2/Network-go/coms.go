package main

import (
	"Network-go/network/bcast"
	"Network-go/network/localip"
	"flag"
	"fmt"
	"os"
)

type msg struct {
	text string
	//number uint
}

type con_status int
type tx_status int
type rx_status int

const (
	Disconnected con_status = iota
	Connecting
	Connected
)

//The communication FSM goes like this:
//Sender: *send* 0->Transmitted
//Reciever: *recieves and sends back* 0->Recieved
//Sender: *recieves echo* Transmitted -> Confirmed (if same)
//Reciever: *recieves confirmation* Recieved -> Confirmed_recieved

//Some "alive" statement every x ms also needs to be incorporated

const (
	Transmitted tx_status = iota
	Confirmed_transmitted
)

const (
	Recieved rx_status = iota
	Confirmed_recieved
)

type node struct {
	id         string //unique id for the network
	connection con_status
	transmit   tx_status
	recieve    rx_status
}

func transmit(tx_ch chan string) {
	//
}

func recieve(ch chan string, port int) {
	bcast.Receiver(port, ch)
}

func main() {
	//Give each one ID
	var id string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.Parse()

	// ... or alternatively, we can use the local IP address.
	// (But since we can run multiple programs on the same PC, we also append the
	//  process ID)
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}

	//Check for active nodes on ip:port
	//peerUpdateCh := make(chan peers.PeerUpdate)
	//go peers.Receiver(30000, peerUpdateCh)

	//define the datatype we want to recive
	rx_ch := make(chan msg)

	//recieve thread
	go bcast.Receiver(30000, rx_ch) //recieve(rx_ch, 30000)

	//printing loop
	for {
		select {
		case in := <-rx_ch:
			println("Received:  %#v\n", in.text)
		}
	}
}
