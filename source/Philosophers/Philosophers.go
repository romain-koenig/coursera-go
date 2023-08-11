package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	index           int
	id              int
	portionsLeft    int
	leftCS, rightCS *ChopS
}

type Channels struct {
	requestChannel        chan int
	personalChannels      []chan bool
	finishedEatingChannel chan int
}

// coinFlip returns true or false randomly. Simple helper function.
func coinFlip() bool {
	return rand.Intn(2) == 1
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
func Host(communicationChannels Channels,
	numberOfPhilosophers int,
	wg *sync.WaitGroup) {
	defer wg.Done()

	// philosophersIsEating is a slice that will keep track of which philosophers are eating.
	philosophersIsEating := make([]bool, numberOfPhilosophers)
	waitingQueue := []int{} // A queue to store waiting philosophers.

	philosophersDone := 0

	for {
		select {

		case finishedPhiloID := <-communicationChannels.finishedEatingChannel:
			philosophersDone++
			philosophersIsEating[finishedPhiloID] = false
			if philosophersDone == numberOfPhilosophers {
				return
			}

		case philoID := <-communicationChannels.requestChannel:
			if !philosophersIsEating[philoID] {
				if howManyPhilosophersAreEating(philosophersIsEating) < 2 {
					// Grant permission
					philosophersIsEating[philoID] = true
					communicationChannels.personalChannels[philoID] <- true
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
					communicationChannels.personalChannels[nextPhilo] <- true
				}
			}
		case <-time.After(time.Second * 10):
			// Some reasonable timeout
			fmt.Println("Timeout, host is exiting.") // Note: this should never happen. It's only a failsafe.
			return
		}
	}
}

func (p Philo) eat(communicationChannels Channels,
	wg *sync.WaitGroup) {

	defer wg.Done()

	for p.portionsLeft > 0 {

		// Ask for permission to eat
		communicationChannels.requestChannel <- p.index

		// Wait for permission

		<-communicationChannels.personalChannels[p.index]

		p.pickCS()

		fmt.Println("starting to eat", p.id)

		// The act of eating is just decreasing the number of portions left, but we can imagine that it's a more complex process.
		p.portionsLeft--

		fmt.Println("finishing eating", p.id)

		p.releaseCS()

		// Inform the host that you're done eating your portion
		communicationChannels.requestChannel <- p.index
	}

	// Once all portions are eaten
	communicationChannels.finishedEatingChannel <- p.index
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

	numberOfPhilosophers := flag.Int("n", 5, "Number of philosophers")
	numberOfPortions := flag.Int("p", 3, "Number of portions per philosopher")
	flag.Parse()

	if *numberOfPhilosophers < 2 {
		fmt.Println("There must be at least two philosophers.")
		return
	}

	if *numberOfPhilosophers > 8000 {
		fmt.Println("There can't be more than 8000 philosophers (hard limit on the number of goroutines).")
		return
	}

	if *numberOfPortions < 1 {
		fmt.Println("There must be at least one portion per philosopher.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the dining philosophers problem!")
	fmt.Println("-------------------------------------------")
	fmt.Println()
	fmt.Println("You can change the number of philosophers and the number of portions per philosopher using the -n and -p flags.")
	fmt.Println("Example: >go run Philosophers.go -n 10 -p 5")
	fmt.Println()
	fmt.Println("Number of philosophers:", *numberOfPhilosophers)
	fmt.Println("Number of portions per philosopher:", *numberOfPortions)
	fmt.Println()

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Add all the philosophers + the host to the WaitGroup
	wg.Add(*numberOfPhilosophers + 1)

	// initialize the ChopSticks
	CSticks := make([]*ChopS, *numberOfPhilosophers)
	for i := 0; i < *numberOfPhilosophers; i++ {
		CSticks[i] = new(ChopS)
	}

	// Initialize the Philosophers
	philos := make([]*Philo, *numberOfPhilosophers)
	for i := 0; i < *numberOfPhilosophers; i++ {
		philos[i] = &Philo{i, i + 1, *numberOfPortions, CSticks[i], CSticks[(i+1)%*numberOfPhilosophers]}
	}

	// Create the request channel
	requestChannel := make(chan int)

	// Create the personal channels
	personalChannels := make([]chan bool, *numberOfPhilosophers)
	for i := 0; i < *numberOfPhilosophers; i++ {
		personalChannels[i] = make(chan bool)
	}

	// Add a channel for philosophers to notify the host they're done
	finishedEatingChannel := make(chan int)

	communicationChannels := Channels{requestChannel, personalChannels, finishedEatingChannel}

	// Start the host

	go Host(communicationChannels, *numberOfPhilosophers, &wg)

	// Make the philosophers eat
	for i := 0; i < *numberOfPhilosophers; i++ {
		go philos[i].eat(communicationChannels, &wg)
	}

	// Maker sure that the host has finished before exiting
	wg.Wait()

	fmt.Println("All philosophers are done eating, host has exited, program is done.")

}
