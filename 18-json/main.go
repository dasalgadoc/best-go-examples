package main

import (
	"encoding/json"
	"fmt"
)

// public struct, public properties for Unmarshal
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func main() {
	representations()

	jsonToStruct()

	structToJson()

	j := JsonPathExample{}
	j.Run()
}

func structToJson() {
	simple := User{
		Name:    "daniela",
		Age:     25,
		Email:   "example@daniela.com",
		IsAdmin: true,
	}

	var data []byte
	data, _ = json.Marshal(simple)
	fmt.Println(string(data))
	fmt.Println()
}

func jsonToStruct() {
	// https://pkg.go.dev/encoding/json
	// https://gobyexample.com/json
	data := []byte(`{"name": "daniela", "age": 25, "is_admin": false}`)

	var daniUser User
	err := json.Unmarshal(data, &daniUser)
	if err != nil {
		fmt.Println("Error:  ", err)
	}

	fmt.Println(daniUser)
	fmt.Println()
}

func representations() {
	data := simpleStruct()
	fmt.Println("struct: ", data)
	fmt.Println()

	slice := structSlice()
	fmt.Println("slice: ", slice)
	fmt.Println()

	sliceOfSlices := sliceOfSlices()
	fmt.Println("slice of slice: ", sliceOfSlices)
	fmt.Println()
}

func sliceOfSlices() [][]User {
	sliceOfSlices := [][]User{}
	for i := 0; i < 5; i++ {
		value := structSlice()
		sliceOfSlices = append(sliceOfSlices, value)
	}
	return sliceOfSlices
}

func structSlice() []User {
	data := []User{}

	for i := 0; i < 10; i++ {
		data = append(data, simpleStruct())
	}
	return data
}

func simpleStruct() User {
	data := User{
		Name:    "Diego",
		Age:     30,
		Email:   "example@example.com",
		IsAdmin: true,
	}

	return data
}
