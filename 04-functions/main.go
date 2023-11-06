package main

import "fmt"

func main() {
	fmt.Println("Executing functions...")
	simpleProcedure()
	procedureWithParams(10, "Diego")
	procedureWithParamsGoodPractice(20, 30)

	fmt.Println(function())
	fmt.Println()
	fmt.Println(biFunction())
	fmt.Println()

	// Ignoring values
	firstVar, _ := biFunction()
	fmt.Println()
	_, secondVar := biFunction()
	fmt.Println(firstVar, secondVar)
	fmt.Println()

	fmt.Println("Add two numbers: ", sumASlice(1, 2))
	fmt.Println("Add three numbers: ", sumASlice(1, 2, 3))
	fmt.Println("Add five numbers: ", sumASlice(1, 2, 3, 4, 5))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Add five numbers: ", sumASlice(slice...))

	fmt.Println()

	// Anonymous functions
	func() {
		fmt.Println("Hi! I'm a anonymous function")
		fmt.Println("End of anonymous function")
		fmt.Println()
	}()

	func(text string) {
		fmt.Println("Hi! I'm a anonymous function")
		fmt.Printf("\t...but with this %s parameter\n", text)
		fmt.Println("End of anonymous function")
		fmt.Println()
	}("the param")

	// Assign function
	myAnonFunc := func() {
		fmt.Printf("Anon\t\n")
	}

	myAnonFunc()
	myAnonFunc()
	myAnonFunc()

	functions := FunctionOnMap{
		name: "TEN",
	}
	functions.Run()
}

func simpleProcedure() {
	fmt.Println("Hi! I'm a simple procedure")
	fmt.Println("I have no return value")
	fmt.Println("... but maybe a side effect")
	fmt.Printf("End of simple procedure\n\n")
}

func procedureWithParams(number int, text string) {
	fmt.Printf("Hi! I receive this params %d - %s\n", number, text)
	fmt.Printf("End of procedure with params\n\n")
}

func procedureWithParamsGoodPractice(firstNumber, secondNumber int) {
	fmt.Printf("Hi! check my signature for see a good practice\n")
	fmt.Printf("End of procedure with params Good Practice\n\n")
}

func function() int {
	fmt.Println("Hi! I'm a function the I'll return...")
	return 100
}

func biFunction() (int, int) {
	fmt.Println("Hi! I'm a function the I'll return (two values)...")
	return 200, 300
}

func sumASlice(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}
