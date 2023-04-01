package main

import "fmt"

func main() {
	arrays()

	slices()

	slicing()

	slicingLoop()
}

func slicingLoop() {
	slice := []string{"hello", "big", "world"}
	fmt.Println(slice)

	for i, value := range slice {
		fmt.Println(i, value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	for i, _ := range slice {
		fmt.Println(i)
	}
}

func slicing() {
	slice := []int{100, 200, 300, 400, 500}
	fmt.Println("First three: ", slice[:3])
	fmt.Println("Last three: ", slice[2:])
	fmt.Println("Mid: ", slice[2:3])
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))

	// deleted
	slice = append(slice[:len(slice)-1])
	fmt.Println(slice)
}

func slices() {
	// Slice
	var slice []int
	fmt.Println("Zero value slice", slice)
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))

	initialSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice initialized", initialSlice)
	fmt.Println("Length: ", len(initialSlice))
	fmt.Println("Capacity: ", cap(initialSlice))

	// Append
	fmt.Println("Appending to slice")
	slice = append(slice, 100)
	fmt.Println(slice)

	slice = append(slice, 200, 300)
	fmt.Println(slice)

	secondSlice := []int{400, 500, 600}
	slice = append(slice, secondSlice...)
	fmt.Println(slice)
}

func arrays() {
	// Array
	var array [5]int
	fmt.Println("Zero value array", array)

	array[1] = 20
	fmt.Println("Assign element", array)
	fmt.Println("Length: ", len(array))
	fmt.Println("Capacity: ", cap(array))

	initializedArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array initialized ", initializedArray)
	fmt.Println("Length: ", len(initializedArray))
	fmt.Println("Capacity: ", cap(initializedArray))
}
