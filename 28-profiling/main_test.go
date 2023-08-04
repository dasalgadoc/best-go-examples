package main

import "testing"

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
