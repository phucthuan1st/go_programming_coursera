package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers   = 5
	maxParallelEating = 2
	maxMeals          = 3
)

var times = []string{"1st", "2nd", "3rd"}

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philosopher) eat(host chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < maxMeals; i++ {
		// Request permission from the host
		<-host

		fmt.Printf("[Philosopher %d] is allowed to eat\n", p.id)

		fmt.Printf("[Philosopher %d] is waiting to pick up [left chopsticks %d]\n", p.id, (p.id))
		p.leftChopstick.Lock()
		fmt.Printf("[Philosopher %d] is picked up [left chopsticks %d]\n", p.id, (p.id))

		fmt.Printf("[Philosopher %d] is waiting to pick up [right chopsticks %d]\n", p.id, (p.id+1)%numPhilosophers)
		p.rightChopstick.Lock()
		fmt.Printf("[Philosopher %d] is picked up [right chopsticks %d]\n", p.id, (p.id+1)%numPhilosophers)

		fmt.Printf("[Philosopher %d] starting to eat for [%v times]\n", p.id, times[i])
		timeToEat := rand.Intn(1000)
		time.Sleep(time.Millisecond * time.Duration(timeToEat)) // Simulating eating
		fmt.Printf("[Philosopher %d] finishing eating for [%v times] after %vms\n", p.id, times[i], timeToEat)

		p.rightChopstick.Unlock()
		fmt.Printf("[Philosopher %d] is put down [right chopsticks %d]\n", p.id, (p.id+1)%numPhilosophers)
		p.leftChopstick.Unlock()
		fmt.Printf("[Philosopher %d] is put down [left chopsticks %d]\n", p.id, (p.id))

		// Release permission to the host
		host <- true
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	host := make(chan bool, maxParallelEating)
	for i := 0; i < maxParallelEating; i++ {
		host <- true
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat(host, &wg)
	}

	wg.Wait()
}
