package main

func mergeSort(unsorted []int) []int {
	// base case
	if len(unsorted) <= 1 {
		return unsorted
	}

	// divide
	mid := len(unsorted) / 2

	// conquer
	left := mergeSort(unsorted[:mid])
	right := mergeSort(unsorted[mid:])

	// combine
	return merge(left, right)
}

func merge(left, right []int) []int {
	// merge two sorted slices
	merged := make([]int, 0)
	leftIndex := 0
	rightIndex := 0

	// compare and merge
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			merged = append(merged, left[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, right[rightIndex])
			rightIndex++
		}
	}

	// append remaining elements
	merged = append(merged, left[leftIndex:]...)
	merged = append(merged, right[rightIndex:]...)

	return merged
}
