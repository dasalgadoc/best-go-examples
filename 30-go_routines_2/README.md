# 👐🏻 Go Routines

See the first [example](../14-go_routines/)

## What is a GoRoutine?

A goroutine is a lightweight thread managed by the Go.
It can be thought as a function executing concurrently through `go` sentence.

__Basic example__:
```go
package main

import (
	"fmt"
	"time"
)

func commonFunc() {
    fmt.Println("I'm a common func")
}

func main() {
    go commonFunc() // This is a goroutine
	
    fmt.Println("I'm the main routine")
    time.Sleep(1 * time.Second)
}
```

## WaitGroups

Waitgroup is a structure that allows to wait for the execution of a group of added goroutines.

- Function `Add` adds the number of goroutines to wait for.
- Function `Done` decrements the counter of goroutines to wait for.
- Function `Wait` blocks the execution until the counter is zero.

__Basic example__:
```go
package main

import (
	"fmt"
	"sync"
)

func commonFunc(wg *sync.WaitGroup) {
    fmt.Println("I'm a common func")
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    
    go commonFunc(&wg)
    
    fmt.Println("I'm the main routine")
    wg.Wait()
}
```

## Channels

Channels are a way to communicate between goroutines.

There is two types of channels:
- Buffered channels: Has a capacity, and it can be filled until it is full.
- Unbuffered channels: Has no capacity, and it can be filled until it is empty.

## Selects

Selects are a way to listen to multiple channels at the same time.

See this [example](../15-channels/)
