package main

import "fmt"

type FunctionOnMap struct {
	functions map[string]func()
	name      string
}

func (f *FunctionOnMap) Run() {
	f.functions = map[string]func(){
		"HELLO": sayHello,
		"TEN":   printTenNumbers,
	}

	f.functions[f.name]()
}

func sayHello() {
	fmt.Println("Hello!!!")
}

func printTenNumbers() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(numbers)
}
