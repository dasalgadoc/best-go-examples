package main

import "fmt"

// Traditional function for the sum.
func add(x, y int) int {
	return x + y
}

// Curry function for the sum.
func curryAdd(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	// Add two numbers.
	result1 := add(3, 5)
	fmt.Println("Result 1:", result1)

	// Currying the function add.
	curriedAdd := curryAdd(3)

	// Call the curried function.
	result2 := curriedAdd(5)
	fmt.Println("Result 2:", result2)

	// Call the curried function.
	result3 := curryAdd(10)(30)
	fmt.Println("Result 3:", result3)
}
