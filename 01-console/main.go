package main

import (
	"fmt"
	"time"
)

func main() {

	// fmt handles with IO: https://pkg.go.dev/fmt

	fmt.Println("Hello World")

	fmt.Printf("Its %s \n", time.Now().String())

	fmt.Println("First word ", "Second word")
}
