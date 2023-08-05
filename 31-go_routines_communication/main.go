package main

import (
	"fmt"
	"math"
	"sync"
)

type FibonacciValueObject struct {
	Value int
}

const NUMBER_PER_ROUTINE = 5

func main() {
	fmt.Println("By shared memory")
	memory := bySharedMemoryNumberFinder(4181)
	fmt.Println("Result: ", memory)

	fmt.Println("By message passing")
	passing := byMessagePassing(75025)
	fmt.Println("Result: ", *passing)
}

func bySharedMemoryNumberFinder(id int) FibonacciValueObject {
	routinesNumber := calculateRoutines(len(FIBONACCI_SERIES_NUMBERS))

	wg := &sync.WaitGroup{}
	wg.Add(routinesNumber)

	var result FibonacciValueObject

	for i := 0; i < routinesNumber; i++ {
		begin := i * NUMBER_PER_ROUTINE
		end := begin + NUMBER_PER_ROUTINE
		go func(number, begin, end int, numbers []int, result *FibonacciValueObject, wg *sync.WaitGroup) {
			for i := begin; i <= end; i++ {
				if i >= len(numbers) {
					continue
				}
				if numbers[i] == number {
					*result = FibonacciValueObject{Value: numbers[i]}
				}
			}
			wg.Done()
		}(id, begin, end, FIBONACCI_SERIES_NUMBERS, &result, wg)
	}

	wg.Wait()
	return result
}

func byMessagePassing(id int) *FibonacciValueObject {
	routinesNumber := calculateRoutines(len(FIBONACCI_SERIES_NUMBERS))

	in := make(chan FibonacciValueObject)
	done := make(chan bool, routinesNumber)

	for i := 0; i < routinesNumber; i++ {
		begin := i * NUMBER_PER_ROUTINE
		end := begin + NUMBER_PER_ROUTINE
		toSearch := make([]int, NUMBER_PER_ROUTINE)
		copy(toSearch[:], FIBONACCI_SERIES_NUMBERS[begin:end])

		go func(toSearch []int, in chan FibonacciValueObject, done chan bool) {
			for _, number := range toSearch {
				if number == id {
					in <- FibonacciValueObject{Value: number}
				}
			}
			done <- true
		}(toSearch, in, done)
	}

	var result FibonacciValueObject
	i := 0
	for i < routinesNumber {
		select {
		case result = <-in:
			return &result
		case <-done:
			i++
		}
	}

	return nil
}

func calculateRoutines(lendOfNumbers int) int {
	return int(math.Ceil(float64(lendOfNumbers) / float64(NUMBER_PER_ROUTINE)))
}
