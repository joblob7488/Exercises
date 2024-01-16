package main

import (
	//"Network-go/network/bcast"
	//"Network-go/network/localip"
	"Network-go/network/peers"
	//"flag"
	//"fmt"
	//"os"
	//"time"
)

type msg struct {
	msg string
	number uint
}

type con_status int
type tx_status int
type rx_status int

const(
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

const(
	Transmitted tx_status = iota
	Confirmed_transmitted
)

const(
	Recieved rx_status = iota
	Confirmed_recieved
)

type node struct{
	id string //unique id for the network
	connection con_status
	transmit tx_status
	recieve rx_status
}

func transmit(tx_ch chan){
	//
}

func recieve(ch chan, port uint){
	peers.Receiver(port, ch)
}

func main(){
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}

	rx := make(rx_ch chan)
	go recieve(rx_ch,30000)
	for{
		println(<-rx)
	}

}