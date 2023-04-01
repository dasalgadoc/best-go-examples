package main

import "fmt"

func main() {

	// Data types: https://www.w3schools.com/go/go_data_types.php

	// Constants
	// Form 1 - Definition with type
	const pi float64 = 3.14
	// Form 2 - Definition explicit
	const au = 1.14

	fmt.Printf("PI value: %s\n", pi)
	fmt.Printf("AU value: %s\n", au)

	// Variables
	// Form 1 - implicit and zero value
	var area int
	// Form 2 - implicit and value
	var height int = 14
	// Form 3 - explicit and value
	base := 12

	area = height * base
	fmt.Printf("Base: %s \nHeight: %s \nArea: %s \n", base, height, area)

	// Zero values:
	var zeroInteger int
	var zeroReal float64
	var zeroString string
	var zeroBoolean bool

	fmt.Println("Integer by default: ", zeroInteger)
	fmt.Println("Real by default: ", zeroReal)
	fmt.Println("String by default: ", zeroString)
	fmt.Println("Boolean by default: ", zeroBoolean)

	// Defining types
	type money int

	var price money
	price = 4000
	fmt.Printf("Set a variable: %T \t %v \n", price, price)

}
