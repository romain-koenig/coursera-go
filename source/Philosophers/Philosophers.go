package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberOfPortions = 3

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func coinFlip() bool {
	coinFlip := rand.Intn(2)
	if coinFlip == 1 {
		return true
	}
	return false
}

func (p Philo) eat() {

	for i := numberOfPortions; i > 0; i-- {

		p.pickCS()

		fmt.Println("starting to eat", p.id)

		p.releaseCS()
	}
}

func (p Philo) releaseCS() {
	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

func (p Philo) pickCS() {
	if coinFlip() {
		p.leftCS.Lock()
		p.rightCS.Lock()
	} else {
		p.rightCS.Lock()
		p.leftCS.Lock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello, Philosophers!")

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Create 5 channels
	chans := make([]chan bool, 5)

	// Add 5 to the WaitGroup
	wg.Add(5)

	// Start the host
	go Host(&wg, chans)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

}

// Host is the goroutine that will handle the two tokens, give them to the philosophers, and take them back when they're done
func Host(wg *sync.WaitGroup, chans []chan bool) {

	defer wg.Done() // When the function is done, tell the WaitGroup

	select {
	case a := <-chans[0]:
		fmt.Println(a)
	case b := <-chans[1]:
		fmt.Println(b)
	case c := <-chans[2]:
		fmt.Println(c)
	case d := <-chans[3]:
		fmt.Println(d)
	case e := <-chans[4]:
		fmt.Println(e)
	default:
		fmt.Println("No Message Received")
	}

}
