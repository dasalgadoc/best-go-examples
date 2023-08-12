package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	numWorkers        = 5
	elementsPerWorker = 20
)

func main() {
	elements := GetPokemonList().Results

	syncMethod(elements)

	asyncMethod(elements, false)
	asyncMethod(elements, true)

}

func syncMethod(elements []Pokemon) {
	now := time.Now()

	log.Println("Sync Execution started")

	for _, element := range elements {
		_ = GetPokemonHeight(element.Name)
	}

	elapsed := time.Since(now)
	log.Printf("Sync Execution took %s", elapsed)
}

func asyncMethod(elements []Pokemon, stepped bool) {
	log.Println("Async Execution started")

	totalElements := len(elements)
	now := time.Now()
	var wg sync.WaitGroup
	ch := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		start := i * (totalElements / numWorkers)
		end := (i + 1) * (totalElements / numWorkers)
		if end > totalElements {
			end = totalElements
		}
		if stepped {
			go steppeWorker(i, elements[start:end], &wg, ch)
		} else {
			go worker(i, start, end, elements, &wg, ch)
		}
	}

	wg.Wait()
	close(ch)

	for range ch {
	}
	fmt.Println("All workers have finished")

	elapsed := time.Since(now)
	log.Printf("Async Execution took %s", elapsed)
}

func worker(id int, start, end int, elements []Pokemon, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := start; i < end; i++ {
		_ = GetPokemonHeight(elements[i].Name)
	}
	ch <- id
}

func steppeWorker(id int, elements []Pokemon, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	semaphore := make(chan Pokemon, elementsPerWorker)
	done := make(chan Pokemon)

	go func() {
		wg := sync.WaitGroup{}

		for _, element := range elements {
			semaphore <- element
			wg.Add(1)

			go func(e Pokemon) {
				defer func() {
					<-semaphore
					wg.Done()
				}()

				_ = GetPokemonHeight(e.Name)
			}(element)
		}

		wg.Wait()
		close(done)
	}()

	<-done
	ch <- id
}
