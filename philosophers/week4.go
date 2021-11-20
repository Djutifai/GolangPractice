package main

import (
	"fmt"
	"sync"
)

type Host struct {
	eatingRightNow int
}

var host = new(Host)

type ChopS struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *ChopS
}

func (p Philosopher) eat(wg *sync.WaitGroup, number int) {
	for {
		for host.eatingRightNow > 2 { // an infinite loop before host will "allow" to enter eating phase
		}
		if number%2 == 0 { //here is the defence of possible deadlock
			p.leftCS.Lock()  // if philosopher's number is even it will grab left chopstick first
			p.rightCS.Lock() // if the number is odd then philosopher will grab the right one
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}
		host.eatingRightNow++
		fmt.Println("starting to eat ", number)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		host.eatingRightNow--
		fmt.Println("finishing eating ", number)
	}
	wg.Done()
}

func main() {
	fmt.Printf("Hello and welcome to my project: Philosophers\nWe are starting\n\n")
	chopsSet := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chopsSet[i] = new(ChopS)
	}
	philoSet := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philoSet[i] = &Philosopher{chopsSet[i], chopsSet[(i+1)%5]}
	}
	var wg = sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philoSet[i].eat(&wg, i+1)
	}
	wg.Wait()
	fmt.Println("All philosophers ate their food three times, the dinner is over :D")
}
