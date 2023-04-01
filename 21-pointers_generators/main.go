package main

import "fmt"

type MyStruct struct {
	Attr string
}

type Flowable interface{}

var ImplementedOne = map[string]Flowable{
	"base": &MyStruct{},
}

var ImplementedTwo = map[string]func() Flowable{
	"base": GetStruct,
}

func GetStruct() Flowable {
	return &MyStruct{}
}

func main() {

	withDirect()

	withGenerator()
}

func withDirect() {
	var flows []Flowable

	for i := 0; i < 5; i++ {
		flows = append(flows, ImplementedOne["base"])
	}

	for _, flow := range flows {
		fmt.Printf("%v - %p\n", flow, flow)
	}
}

func withGenerator() {
	fmt.Println("With Generator")

	var flows []Flowable

	for i := 0; i < 5; i++ {
		flows = append(flows, ImplementedTwo["base"]())
	}

	for _, flow := range flows {
		fmt.Printf("%v - %p\n", flow, flow)
	}
}
