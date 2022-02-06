package main

import (
	"fmt"
	"sync"
	"time"
)

var mt sync.Mutex
var wg sync.WaitGroup
var ch chan bool
var eaten int

func main() {
	/*ch = make(chan bool, 2)
	wg.Add(2)
	for i := 1; i <= 5; i++ {
		fmt.Println("Creating philosopher", i)
		go philosopher(i)
	}

	select {
	case <-ch:
		wg.Done()
	default:
		fmt.Println("Waiting for philosophers to eat")
	}*/
	mt.Lock()
	go host()
	for i := 1; i <= 5; i++ {
		go philosopher(i)
	}

	<-ch
	fmt.Println("All philosophers have eaten")
	select {}
}

func philosopher(num int) {
	//fmt.Println("Philosopher", num, "is eating")
	fmt.Println("Philosopher", num, "started")
	for i := 0; i < 3; i++ {
		mt.Lock()
		time.Sleep(1 * time.Second)
		fmt.Println("Philosopher", num, "is eating")
	}
}

func host() {
	for {
		time.Sleep(2 * time.Second)
		eaten++
		fmt.Println("Host loop:", eaten)
		mt.Unlock()

		if eaten == 15 {
			mt.Unlock()
			ch <- true
		}
	}
}
