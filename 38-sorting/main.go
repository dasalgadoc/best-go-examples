package main

import (
	"fmt"
	"math/rand"
)

func main() {
	standardExamples()
}

func standardExamples() {
	unsorted := buildRandomIntSlice()
	fmt.Println("Unsorted slice", unsorted)
	ascStdSort(unsorted)

	unsorted = buildRandomIntSlice()
	fmt.Println("Unsorted slice", unsorted)
	desStdSort(unsorted)

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println("Unsorted slice", people)
	ascStdSortPeople(people)

	sortedOdds := buildOddIntSlice()
	fmt.Println("Sorted odds", sortedOdds)
	target := 5
	index := binarySearchStd(sortedOdds, target)
	fmt.Println("Index of 5 in sorted odds (Search)", index, sortedOdds[index] == target)

	target = 14
	index = binarySearchStd(sortedOdds, target)
	fmt.Println("Index of 14 in sorted odds (Search)", index, sortedOdds[index] == target)

	i, f := binaryFindStd(sortedOdds, 5)
	fmt.Println("Index of 5 in sorted odds (Find?)", i, f)

	i, f = binaryFindStd(sortedOdds, 14)
	fmt.Println("Index of 14 in sorted odds (Find?)", i, f)

	sortPlanetsExample()
}

func buildRandomIntSlice() []int {
	// Make a slice of 20 random integers
	size := 20
	upperbound := 1000
	unsorted := make([]int, 0)
	for i := 0; i < size; i++ {
		unsorted = append(unsorted, rand.Intn(upperbound))
	}
	return unsorted
}

func buildOddIntSlice() []int {
	// Make a slice with first 20 odd integers
	size := 20
	sortedOdds := make([]int, 0)
	for i := 1; i < size*2; i += 2 {
		sortedOdds = append(sortedOdds, i)
	}
	return sortedOdds
}
