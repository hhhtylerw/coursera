package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Philosopher struct {
	Name           int
	LeftChopStick  *Chopsticks
	RightChopStick *Chopsticks
}

type Chopsticks struct {
	sync.Mutex
}

func main() {
	philosopherCount := 5
	chopsticks := make([]*Chopsticks, philosopherCount)
	for i := 1; i <= philosopherCount; i++ {
		chopsticks[i-1] = new(Chopsticks)
	}

	philosophers := make([]*Philosopher, philosopherCount)
	for i := 1; i <= philosopherCount; i++ {
		philosophers[i-1] = &Philosopher{i, chopsticks[i-1], chopsticks[(i)%philosopherCount]}
	}

	for _, p := range philosophers {
		wg.Add(1)
		go p.eat()
	}
	wg.Wait()
}

func (p Philosopher) eat() {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		p.LeftChopStick.Lock()
		p.RightChopStick.Lock()
		fmt.Println(p.Name, "is starting to eat.")
		p.LeftChopStick.Unlock()
		p.RightChopStick.Unlock()
		fmt.Println(p.Name, "has finished eating.")
		time.Sleep(1 * time.Second)
	}
}
