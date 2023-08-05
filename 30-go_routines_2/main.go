package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Buffered channel...")
	bufferedChannel()
	fmt.Println("...Done")

	time.Sleep(1 * time.Second)

	fmt.Println("Unbuffered channel...")
	unbufferedChannel()
	fmt.Println("...Done")

	time.Sleep(3 * time.Second)
	fmt.Println("Selects...")
	selects()
	fmt.Println("...Done")
}

func unbufferedChannel() {
	ch := make(chan int) // unbuffered channel
	from, to := 1, 10
	go genNumbers(from, to, ch)
	for i := from; i <= to; i++ {
		fmt.Printf("Received: %d\n", <-ch)
	}
}

func bufferedChannel() {
	ch := make(chan int, 10) // buffered channel
	from, to := 1, 10
	go genNumbers(from, to, ch)
	for i := from; i <= to; i++ {
		fmt.Printf("Received: %d\n", <-ch)
	}
}

func selects() {
	from, to := 1, 10
	genCh := make(chan int)
	timeoutCh := make(chan bool)

	go genNumbers(from, to, genCh)
	go timeout(1, timeoutCh)

	for {
		select {
		case n := <-genCh:
			fmt.Printf("Received: %d\n", n)
		case <-timeoutCh:
			fmt.Println("Timeout")
			return
		}
	}

}

func timeout(n int, ch chan bool) {
	time.Sleep(time.Duration(n) * time.Second)
	ch <- true
}

func genNumbers(from, to int, ch chan int) {
	for i := from; i <= to; i++ {
		ch <- i
		fmt.Printf("Published: %d\n", i)
	}
}
