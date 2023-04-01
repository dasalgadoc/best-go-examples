package main

import "fmt"

func main() {
	var odd int = 33
	var even int = 20

	if odd%2 != 0 {
		fmt.Println("odd number")
	}

	if odd%2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	if odd%2 == 0 {
		fmt.Println("even number")
	} else if even%2 == 0 {
		fmt.Println("odd number")
	} else {
		fmt.Println("border case")
	}

	fmt.Println()

	var mult int = odd * even
	var module bool = mult%2 == 0

	switch module {
	case true:
		fmt.Println("even multiplication")
	default:
		fmt.Println("odd multiplication")
	}

	switch {
	case module && even > odd:
		fmt.Println("even mult and even greater than odd")
	case module && even < odd:
		fmt.Println("even mult and odd greater than even")
	case module && even == odd:
		fmt.Println("even mult and equal numbers")
	default:
		fmt.Println("odd mult")
	}

}
