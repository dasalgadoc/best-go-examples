package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	waitGroup  sync.WaitGroup
	cap        int
	iterations int
}

func NewSemaphore(cap, iterations int) Semaphore {
	return Semaphore{
		cap:        cap,
		iterations: iterations,
	}
}

func (s *Semaphore) Run() {
	// limited channel. Cap =
	channel := make(chan int, s.cap)

	// use the channel until it is filled
	for i := 0; i < s.iterations; i++ {
		s.waitGroup.Add(1)
		channel <- 1

		// only can execute cap times
		go doSomething(i, &s.waitGroup, channel)
	}

}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("id: %d started \n", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("id: %d finished \n", i)

	// free a space
	<-c
}
