package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	host Host
	wg   sync.WaitGroup
	once sync.Once
)

type Host struct {
	// slots is slots for philosophers to eat concurrently
	slots chan int
}

func (h *Host) Init() {
	// Two slots for two philosophers to eat concurrently
	// by this way, host will automatically block other
	// philosophers when slots is full
	h.slots = make(chan int, 2)
}

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id         int
	lChopstick *Chopstick
	rChopstick *Chopstick
}

func (p *Philosopher) getPermissionFromHost() {
	// Get permission by assigning his id to slots channel
	host.slots <- p.id
}

func (p *Philosopher) givePermissionBack() {
	// Give slot back for others
	<-host.slots
}

func (p *Philosopher) Eat() {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		p.getPermissionFromHost()

		p.lChopstick.Lock()
		p.rChopstick.Lock()
		fmt.Printf("starting to eat %d\n", p.id)

		// Time duration for eating
		time.Sleep(20 * time.Millisecond)

		fmt.Printf("finishing eating %d\n", p.id)
		p.lChopstick.Unlock()
		p.rChopstick.Unlock()

		p.givePermissionBack()
	}
}

func main() {
	chopstick := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopstick[i] = new(Chopstick)
	}

	philosopher := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosopher[i] =
			&Philosopher{id: i + 1, lChopstick: chopstick[i], rChopstick: chopstick[(i+1)%5]}
	}

	host.Init()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosopher[i].Eat()
	}
	wg.Wait()
}
