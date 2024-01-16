// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing(ch, fin chan bool) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		ch <- true
	}
	println("fin inc")
	fin <- true
}

func decrementing(ch, fin chan bool) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000001; j++ {
		ch <- true
	}
	println("fin dec")
	fin <- true
}

func server(ch_inc, ch_dec chan bool) {
	//TODO: implement the server that will serve the channels
	for {
		select {
		case <-ch_dec:
			i -= 1
		case <-ch_inc:
			i += 1
		}
	}

}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	ch_inc := make(chan bool)
	ch_dec := make(chan bool)
	ch_fin := make(chan bool)

	go server(ch_inc, ch_dec)
	// TODO: Spawn both functions as goroutines
	go incrementing(ch_inc, ch_fin)
	go decrementing(ch_dec, ch_fin)

	for {
		if <-ch_fin && <-ch_fin {
			// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
			// We will do it properly with channels soon. For now: Sleep.
			time.Sleep(500 * time.Millisecond)
			Println("The magic number is:", i)
			break
		}

	}
}
