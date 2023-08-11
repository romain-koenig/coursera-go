package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberOfPortions = 3
const numberOfPhilosophers = 5

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	portionsLeft    int
	leftCS, rightCS *ChopS
}

// coinFlip returns true or false randomly. Simple helper function.
func coinFlip() bool {
	coinFlip := rand.Intn(2)
	if coinFlip == 1 {
		return true
	}
	return false
}

func howManyPhilosophersAreEating(philosophersIsEating []bool) int {
	eatingPhilosophers := 0
	for i := 0; i < len(philosophersIsEating); i++ {
		if philosophersIsEating[i] {
			eatingPhilosophers++
		}
	}
	return eatingPhilosophers
}

// Host is the goroutine that will make sure that only two philosophers are eating at the same time
func Host(requestChannel <-chan int, personalChannels []chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// philosophersIsEating is a slice that will keep track of which philosophers are eating.
	philosophersIsEating := make([]bool, numberOfPhilosophers)
	waitingQueue := []int{} // A queue to store waiting philosophers.

	for {
		select {
		case philoID := <-requestChannel:
			if !philosophersIsEating[philoID] {
				if howManyPhilosophersAreEating(philosophersIsEating) < 2 {
					// Grant permission
					philosophersIsEating[philoID] = true
					personalChannels[philoID] <- true
				} else {
					// Add philosopher to waiting queue.
					waitingQueue = append(waitingQueue, philoID)
				}
			} else {
				// Philosopher is done eating.
				philosophersIsEating[philoID] = false

				// Check if there's any philosopher in the queue and grant permission.
				if len(waitingQueue) > 0 {
					nextPhilo := waitingQueue[0]    // Get the next philosopher from the queue.
					waitingQueue = waitingQueue[1:] // Dequeue.

					// Grant permission.
					philosophersIsEating[nextPhilo] = true
					personalChannels[nextPhilo] <- true
				}
			}
		case <-time.After(time.Second * 10):
			// Some reasonable timeout
			fmt.Println("No more requests, host is done.")
			return
		}
	}
}

func (p Philo) eat(requestChannel chan<- int, personalChannels []chan bool) {

	for p.portionsLeft > 0 {

		// Ask for permission to eat
		requestChannel <- p.id

		// Wait for permission

		<-personalChannels[p.id]

		p.pickCS()

		fmt.Println("starting to eat", p.id)

		// The act of eating is just decreasing the number of portions left, but we can imagine that it's a more complex process.
		p.portionsLeft--

		fmt.Println("finishing eating", p.id)

		p.releaseCS()

		// Inform the host that you're done eating your portion
		requestChannel <- p.id
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

	// Add 5 to the WaitGroup
	wg.Add(1)

	// initialize the ChopSticks
	CSticks := make([]*ChopS, numberOfPhilosophers)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Initialize the Philosophers
	philos := make([]*Philo, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		philos[i] = &Philo{i, numberOfPortions, CSticks[i], CSticks[(i+1)%numberOfPhilosophers]}
	}

	// Create the request channel
	requestChannel := make(chan int)

	// Create the personal channels
	personalChannels := make([]chan bool, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		personalChannels[i] = make(chan bool)
	}

	// Start the host

	go Host(requestChannel, personalChannels, &wg)

	// Make the philosophers eat
	for i := 0; i < numberOfPhilosophers; i++ {
		go philos[i].eat(requestChannel, personalChannels)
	}

	// Maker sure that the host has finished before exiting
	wg.Wait()

}
