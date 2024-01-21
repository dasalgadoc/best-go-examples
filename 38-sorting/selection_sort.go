package main

// selectionSort sorting algorithm that divides the input list into two parts:
// the sublist of items already sorted, and the sublist of items remaining to be sorted that occupy the rest of the list.
// Initially, the sorted sublist is empty and the unsorted sublist is the entire input list.
// Each iteration, the algorithm finds the smallest element in the unsorted sublist, and moves it to the end of the sorted sublist.
// It takes O(n^2) time to sort an array of n elements, also takes O(1) space
func selectionSort(unsorted []int) []int {
	for i := 0; i < len(unsorted); i++ {
		minIndex := i
		for j := i + 1; j < len(unsorted); j++ {
			if unsorted[j] < unsorted[minIndex] {
				minIndex = j
			}
		}
		// swap the minimum element with the ith element, 1 to i is sorted in each iteration
		unsorted[i], unsorted[minIndex] = unsorted[minIndex], unsorted[i]
	}

	return unsorted
}

// immutableSelectionSort is an immutable version of selectionSort algorithm that returns a new sorted slice
// in order, slice is passed by reference, a new slice is created to avoid mutating the original slice
func immutableSelectionSort(unsorted []int) []int {
	sorted := make([]int, len(unsorted))
	copy(sorted, unsorted)

	return selectionSort(sorted)
}
