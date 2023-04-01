package main

import "fmt"

func main() {
	// Defer doc: https://go.dev/tour/flowcontrol/12
	defer fmt.Println("...it doesn't even matter")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	var counter int = 10
	for counter < 20 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println()
	for i := counter; i > 0; i-- {
		fmt.Println(i)
	}

	/* Infinite loop
	for {
		fmt.Println("I wont stop! please HELP!")
	}
	*/

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Omitted 2")
			continue
		} else if i == 8 {
			fmt.Println("Stop at 8")
			break
		}
	}

	fmt.Println("In the end")
}
