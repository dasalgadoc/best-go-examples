package main

import "testing"

/* Previous test

func TestFibonacci(t *testing.T) {
	cases := map[string]struct {
		input    int
		expected int
	}{
		"0": {input: 0, expected: 0},
		"1": {input: 1, expected: 1},
		"2": {input: 8, expected: 21},
		"3": {input: 50, expected: 12586269025},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := Fibonacci(tc.input); got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}

func TestFibonacciRefactor(t *testing.T) {
	cases := map[string]struct {
		input    int
		expected int
	}{
		"0": {input: 0, expected: 0},
		"1": {input: 1, expected: 1},
		"2": {input: 8, expected: 21},
		"3": {input: 50, expected: 12586269025},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := FibonacciRefactor(tc.input); got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}

*/

func BenchmarkFibonacci(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//Fibonacci(50)
		FibonacciRefactor(50)
	}
}
