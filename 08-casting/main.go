package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	stringToInt()

	intToString()

	err := stringToBool()

	stringToFloat(err)

	forceToFailCasting(err)
}

func forceToFailCasting(err error) {
	// String to int, force to fail
	fakeIntString := "twenty two"
	fakeStringInt, err := strconv.ParseInt(fakeIntString, 10, 64)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("From %T to %T value = %v", fakeIntString, fakeStringInt, fakeStringInt)
	}
}

func stringToFloat(err error) {
	// String to Float
	floatString := "3.141516"
	stringFloat, err := strconv.ParseFloat(floatString, 64)
	fmt.Printf("From %T to %T value = %v\n", floatString, stringFloat, stringFloat)
}

func stringToBool() error {
	// String to Bool
	booleanString := "true"
	stringBool, err := strconv.ParseBool(booleanString)
	fmt.Printf("From %T to %T value = %v\n", booleanString, stringBool, stringBool)
	return err
}

func intToString() {
	// Int to String
	normalInt := -42
	intString := strconv.Itoa(normalInt)
	fmt.Printf("From %T to %T value = %s\n", normalInt, intString, intString)
}

func stringToInt() {
	// String to Int
	intString := "-42"
	stringInt, err := strconv.Atoi(intString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("From %T to %T value = %d\n", intString, stringInt, stringInt)
}
