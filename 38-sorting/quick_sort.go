package main

// quickSort algorithm relies on recursion and divide and conquer strategy.
// It picks a pivot element and partitions the array around the pivot. Each iteration will move the pivot to its correct position.
// The pivot will be in its correct position when all elements smaller than it are on its left and all elements greater than it are on its right.
// It takes O(n log n) time to sort an array of n elements.
func quickSort(unsorted []int) []int {
	// base case when slice has one element or is empty return itself.
	if len(unsorted) <= 1 {
		return unsorted
	}

	// pick a pivot element as the first element
	pivot := unsorted[0]
	var left, right, sorted []int

	// partition the array around the pivot, left will contain all elements smaller than the pivot, right will contain all elements greater than the pivot.
	for _, v := range unsorted[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// recursively sort the left and right partitions
	sorted = append(sorted, quickSort(left)...)
	// append the pivot to the sorted left partition
	sorted = append(sorted, pivot)
	// append the sorted right partition
	sorted = append(sorted, quickSort(right)...)

	return sorted
}
