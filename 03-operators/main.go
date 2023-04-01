package main

import "fmt"

func main() {
	var a int = 15
	var b int = 10
	var result int

	// Adding
	result = a + b
	fmt.Println("a + b =", result)

	// Subtract
	result = a - b
	fmt.Println("a - b =", result)

	// Multiply
	result = a * b
	fmt.Println("a * b =", result)

	// Division
	result = a / b
	fmt.Println("a / b =", result)

	// Module
	result = a % b
	fmt.Println("a mod b =", result)

	// Incremental
	var counter int = 5
	fmt.Println("counter =", counter)

	counter++
	fmt.Println("counter =", counter)

	// Decremental
	counter--
	fmt.Println("counter =", counter)

}
