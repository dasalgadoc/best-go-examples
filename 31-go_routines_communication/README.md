# 🗣️ Communication Models

Goroutines can communicate and synchronize with each other.

There are 2 ways to communicate between goroutines:
- **Shared memory** (not recommended)
- Message passing

## Shared memory

Shared memory makes use of the same memory space (variables) to communicate between goroutines. Here is when pointers and mutexes come into play.

This model is not recommended because it is sensible to race conditions.

__Example__:

Imagine a function to search an element in a large slice of integers. 

The function will be executed in parallel by multiple goroutines. Each goroutine will search for a different subset of number. 

```go
// Divide the slice in equal parts for each goroutine
func calculateRoutines(lendOfNumbers int) int {
    return int(math.Ceil(float64(lendOfNumbers) / float64(NUMBER_PER_ROUTINE))) 
}

func bySharedMemoryNumberFinder(id int) FibonacciValueObject {
    routinesNumber := calculateRoutines(len(FIBONACCI_SERIES_NUMBERS))
    
    // Wait for all goroutines to finish
    wg := &sync.WaitGroup{}
    wg.Add(routinesNumber)
    
    // This is the shared memory space
    var result FibonacciValueObject
    
    for i := 0; i < routinesNumber; i++ {
        // each goroutine will search for a different subset of numbers
        begin := i * NUMBER_PER_ROUTINE
        end := begin + NUMBER_PER_ROUTINE
		
        // Goroutine: search id in the subset of numbers limited by begin and end, shared memory space is result and wg
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
```

## Message passing

Message passing sends a copy of the messages by using channels. This is the recommended way to communicate between goroutines.

```go
func byMessagePassing(id int) *FibonacciValueObject {
    routinesNumber := calculateRoutines(len(FIBONACCI_SERIES_NUMBERS))
    
	// Channels to send messages between goroutines and mark when they are done
    in := make(chan FibonacciValueObject)
    done := make(chan bool, routinesNumber)
    
    for i := 0; i < routinesNumber; i++ {
		// Make a copy of the slice to search in a subset of numbers
    	begin := i * NUMBER_PER_ROUTINE
    	end := begin + NUMBER_PER_ROUTINE
    	toSearch := make([]int, NUMBER_PER_ROUTINE)
    	copy(toSearch[:], FIBONACCI_SERIES_NUMBERS[begin:end])
        
		// Each goroutine will search for a different subset of numbers
    	go func(toSearch []int, in chan FibonacciValueObject, done chan bool) {
            for _, number := range toSearch {
                if number == id {
					// When number is found, send it to the channel
                    in <- FibonacciValueObject{Value: number}
                }
            }
            done <- true
    	}(toSearch, in, done)
    }
    
    var result FibonacciValueObject
    i := 0
	
	// Search in channels for the first result
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
```
