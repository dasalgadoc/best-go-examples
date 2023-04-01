package main

import "fmt"

type computer struct {
	ram   int
	disk  int
	brand string
}

func (c computer) ping() {
	fmt.Println(c.brand)
}

func (c *computer) duplicateRAM() {
	c.ram *= 2
}

func (c *computer) setRAM(newRam int) {
	c.ram = newRam
}

func main() {
	pointers()

	myPC := computer{ram: 16,
		disk:  512,
		brand: "Asus"}
	fmt.Println(myPC)

	myPC.ping()
	myPC.duplicateRAM()
	fmt.Println(myPC)

	myPC.setRAM(8)
	fmt.Println(myPC)

}

func pointers() {
	// pointer: https://go.dev/tour/moretypes/1
	variable := 100
	pointer := &variable

	var myVar string = "hello world"
	var myPointer *string

	myPointer = &myVar

	fmt.Printf("Variable: %d\tPointer: %v\n", *pointer, pointer)
	fmt.Printf("Variable: %s\tPointer: %v\n", *myPointer, myPointer)

	*pointer = 1000
	fmt.Printf("Variable: %d\tPointer: %v\n", *pointer, pointer)
}
