package main

import "fmt"

func main() {
	elements := [][]string{
		{"DAILY", "NIGHTLY"},
		{"CC"},
		{"CUSTOM"},
		{"TRANSFER"},
	}

	recursiveMethod(elements)

	loopMethod(elements)
}

func loopMethod(elements [][]string) {
	var results [][]string
	index := make([]int, len(elements))

	for {
		var set []string
		for i, e := range elements {
			set = append(set, e[index[i]])
		}
		results = append(results, set)

		for i := len(index) - 1; i >= 0; i-- {
			index[i]++
			if index[i] >= len(elements[i]) {
				index[i] = 0
				if i == 0 {
					break
				}
			} else {
				break
			}
		}

		stop := true
		for i := range index {
			if index[i] != 0 {
				stop = false
				break
			}
		}
		if stop {
			break
		}
	}
	fmt.Println(results)
}

func recursiveMethod(elements [][]string) {
	var results [][]string
	concat(elements, 0, []string{}, &results)
	fmt.Println(results)
}

func concat(elements [][]string, n int, set []string, r *[][]string) {
	// base case
	if n == len(elements) {
		*r = append(*r, set)
		return
	}

	// recursive
	for _, e := range elements[n] {
		c := make([]string, len(set))
		copy(c, set)
		concat(elements, n+1, append(c, e), r)
	}
}
