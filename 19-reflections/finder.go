package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Finder struct {
	person Human
}

type Human struct {
	name     string
	measures Measures
}

type Measures struct {
	height float64
	weight float64
}

func (v *Finder) Run() {
	fmt.Println("Finder example: ")
	measures := Measures{
		height: 1.70,
		weight: 75.0,
	}

	person := Human{
		name:     "Diego",
		measures: measures,
	}

	path := "measures"
	fieldValue := getFieldValue(person, path)

	fmt.Println(fieldValue)
	if !fieldValue.IsValid() {
		fmt.Println("No valid")
	}

}

func getFieldValue(data interface{}, path string) reflect.Value {
	fields := strings.Split(path, ".")
	value := reflect.ValueOf(data)
	for _, field := range fields {
		value = value.FieldByName(field)
	}

	return value
}
