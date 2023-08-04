package main

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciRefactor(n int) int {
	if n <= 1 {
		return n
	}
	prevInt := 0
	nextInt := 1
	for i := 2; i <= n; i++ {
		tempInt := nextInt
		nextInt = prevInt + nextInt
		prevInt = tempInt
	}
	return nextInt
}
