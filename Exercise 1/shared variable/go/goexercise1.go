package main

import (
	"time"
)

type car struct {
	owner, brand       string
	model              uint
	prev_car, next_car *car
}

var (
	Len_lst uint = 3
	Car1         = car{owner: "Heimvik", brand: "Toyota", model: 2005}
	Car2         = car{owner: "Haaland", brand: "Nissan", model: 2006}
	Car3         = car{owner: "Nerheim", brand: "Mitchubishi", model: 2002}
)

func fsm(as, bs int) string {
	text := ""
	for i := 0; i < as; i++ {
		for j := 0; j < bs; j++ {
			if i%2 == 0 {
				text += "moren din"
			} else {
				text += "faren din"
			}
		}
		text += "a"
	}
	return text
}

func cases(one, two, three int) string {
	//check if one is even, two is odd and thrre is larger than 15
	var retval = "Did not get to the end"
	var state = 0
	for i := 1; i < 10; i++ {
		switch state {
		case 0:
			if one%2 == 0 {
				//one is even, next state
				state += 1
			} else {
				break
			}

		case 1:
			if two%2 != 0 {
				//two is odd, next state
				state += 1
			} else {
				break
			}
		case 2:
			if three > 15 {
				//three larger than 15, next state
				state += 1
			} else {
				break
			}
		case 3:
			retval += "Did get to the end"
			break
		}
	}
	return retval
}

func def(a, b int) int {
	defer println("C")
	println("S")
	return a
}

func pointers() int {
	i := 42
	var p *int

	p = &i
	i++
	println(i)
	println(*p)
	return i
}

func list_cars() {
	var current_car_ptr *car = &Car1
	for i := 0; i < 3; i++ {
		println("Owner: ", current_car_ptr.owner, ", Model: ", current_car_ptr.model)
		current_car_ptr = current_car_ptr.next_car

		if current_car_ptr == nil {
			break
		}
	}
}

func send(channel chan int) {
	for i := 0; i < 3; i++ {
		channel <- i
	}
}

func recieve(channels []chan int) {
	for {
		for i := 0; i < 3; i++ {
			select {
			case data := <-channels[i]:
				println("Channel:", i, "wrote: ", data)
			}
		}
	}
}

//Goroutines
func transmission() {
	channels := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < 3; i++ {
		go send(channels[i])
	}
	time.Sleep(time.Second)
	go recieve(channels)
}

func main() {
	//Println("Hello World")
	//Println(fsm(2,3))
	//println(cases(2, 1, 16))
	//println(def(2, 3))

	/* 	Car1.next_car = &Car2
	   	Car2.prev_car = &Car1
	   	Car2.next_car = &Car3
	   	Car3.prev_car = &Car2

	   	list_cars() */

	transmission()
	time.Sleep(time.Second)

}
