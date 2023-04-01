package main

import "fmt"

func main() {
	printHello(0, 10)
}

func printHello(numberToPrint, iteration int) {
	fmt.Println("hello ", numberToPrint)
	if numberToPrint < iteration {
		printHello(numberToPrint+1, iteration)
	}
}
