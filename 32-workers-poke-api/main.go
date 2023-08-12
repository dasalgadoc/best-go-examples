package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	numWorkers = 5
)

func main() {
	elements := GetPokemonList().Results

	syncMethod(elements)

	asyncMethod(elements)
}

func syncMethod(elements []Pokemon) {
	now := time.Now()

	log.Println("Sync Execution started")

	for _, element := range elements {
		_ = GetPokemonHeight(element.Name)
		//log.Println("Pokemon", element.Name, "height is", pokemonStats.Height)
	}

	elapsed := time.Since(now)
	log.Printf("Sync Execution took %s", elapsed)
}

func asyncMethod(elements []Pokemon) {
	log.Println("Sync Execution started")

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
		go worker(i, start, end, elements, &wg, ch)
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
		//log.Println("Pokemon", elements[i].Name, "height is", pokemonStats.Height)
	}
	ch <- id
}
