package main

import (
	"net"
	"os/exec"
	"time"
)

type state int

const (
	master state = iota
	slave state
)

func receiveFromMaster(var port string,chan receivedUpdates byte){
	adr := net.ResolveUDPAddr("udp", port)
	listener := net.ListenUDP("udp", adr)
	defer listener.Close()

	buffer := make([]byte,1024)
	for{
		n := listener.ReadFromUDP(buffer)
		receivedUpdates = buffer[:n]			
	}
}

func sendToSlave(var port string,chan sendUpdates byte){
	adr := net.ResolveUDPAddr("udp",port)
	sender := net.DialUDP("udp",nil,adr)
	defer sender.Close()

	for{
		sender.Write(<-sendUpdates)
	}
}

func masterIsAlive(chan updates byte,var previousUpdate byte){
	select{
	case newUpdate := <-updates:
		if newUpdate != previousUpdate{
			return false
		}else{
			return true
		}
	}
}

func main() {
	//Start the other process
	port := ":20013"
	cmd := exec.Command("gnome-terminal", "--", "go", "run", "process_B.go")
	var currentstate state = slave
	var i int = 0;
	for{
		switch currentstate{
		case master:
			i+= 1
			time.Sleep(1000*time.Milliseconds)
			go sendToSlave(port,i)
		case slave:
			recentUpdate = make(chan byte)
			oldUpdate = <- recentUpdate
			go receiveFromMaster(port,recentUpdates)
			
			if !masterIsAlive(recentUpdates,oldUpdate){
				currentstate = master
				go cmd.Run()
				break
			}
			fmt.Println("Master said:",recentUpdate)
		}
	}
}
