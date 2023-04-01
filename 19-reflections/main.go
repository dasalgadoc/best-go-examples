package main

import (
	"fmt"
	"reflect"
)

type ClientResponse map[string]interface{}

type Person struct {
	Name   string
	Age    int
	Gender string
	Single bool
}

type DynamicParam struct {
	param string
	value string
}

func main() {
	basicReflection()

	dynamicParams()

	finder := Finder{}
	finder.Run()
}

func dynamicParams() {
	site := DynamicParam{param: "site_id", value: "brazil"}
	clientOne := NewClientResponse(site)
	fmt.Println(clientOne)

	person := Person{Name: "Diego", Age: 30}
	clientTwo := NewClientResponse(person)
	fmt.Println(clientTwo)
	fmt.Println()
}

func NewClientResponse(input interface{}) ClientResponse {
	types := reflect.TypeOf(input)
	values := reflect.ValueOf(input)
	response := ClientResponse{}

	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		fmt.Println(field.Name)
		response[field.Name] = values.Field(i)
	}

	return response
}

func basicReflection() {
	// https://go.dev/blog/laws-of-reflection
	// https://www.geeksforgeeks.org/reflection-in-golang/
	person := Person{
		Name:   "John",
		Gender: "male",
		Age:    17,
		Single: true,
	}

	values := reflect.ValueOf(person)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}
}
