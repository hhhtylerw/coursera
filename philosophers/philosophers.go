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
	mt.Lock()
	go host()
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go philosopher(i)
	}

	wg.Wait()
	fmt.Println("All philosophers have eaten")
}

func philosopher(num int) {
	//fmt.Println("Philosopher", num, "is eating")
	fmt.Println("Philosopher", num, "started")
	for i := 0; i < 3; i++ {
		mt.Lock()
		time.Sleep(1 * time.Second)
		fmt.Println("Philosopher", num, "is eating")
	}
	wg.Done()
}

func host() {
	for {
		time.Sleep(2 * time.Second)
		eaten++
		fmt.Println("Host loop:", eaten)
		mt.Unlock()

		if eaten == 15 {
			//mt.Unlock()
			//ch <- true
			return
		}
	}
}
