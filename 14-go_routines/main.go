package main

import (
	"fmt"
	"sync"
)

func main() {
	// Wait group: https://gobyexample.com/waitgroups
	basicWaitGroup()

	semaphore := NewSemaphore(5, 10)
	semaphore.Run()

}

func basicWaitGroup() {
	var waitGroup sync.WaitGroup

	fmt.Println("hello")

	// Add
	waitGroup.Add(1)
	go say("world", &waitGroup)

	waitGroup.Wait()
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
