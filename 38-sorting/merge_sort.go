package main

// mergeSort is a recursive algorithm that continually splits a slice in half call mergeSort on each half recursively.
// It then merges the two sorted halves back together.
// It takes O(n log n) time to sort an array of n elements, also takes O(n) space
func mergeSort(unsorted []int) []int {
	// base case
	if len(unsorted) <= 1 {
		return unsorted
	}

	// divide: Find the midpoint of the slice and divide into sub-slices
	// it takes O(1) time to find the midpoint for iteration but O(log n) in overall
	mid := len(unsorted) / 2

	// conquer: recursively sort the sub-slices
	left := mergeSort(unsorted[:mid])
	right := mergeSort(unsorted[mid:])

	// combine: merge the two sorted sub-slices
	// it takes O(n) time to merge two sorted sub-slices
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

// mergeSortIterative is an iterative version of mergeSort algorithm
func mergeSortIterative(slice []int) []int {
	length := len(slice)

	for size := 1; size < length; size *= 2 {
		for left := 0; left < length-1; left += 2 * size {
			mid := minInt(left+size-1, length-1)
			right := minInt(left+2*size-1, length-1)

			mergeIt(slice, left, mid, right)
		}
	}

	return slice
}

func mergeIt(slice []int, left, mid, right int) {
	leftIndex := mid - left + 1
	rightIndex := right - mid

	leftArray := make([]int, leftIndex)
	rightArray := make([]int, rightIndex)

	for i := 0; i < leftIndex; i++ {
		leftArray[i] = slice[left+i]
	}
	for j := 0; j < rightIndex; j++ {
		rightArray[j] = slice[mid+1+j]
	}

	i, j, k := 0, 0, left
	for i < leftIndex && j < rightIndex {
		if leftArray[i] <= rightArray[j] {
			slice[k] = leftArray[i]
			i++
		} else {
			slice[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < leftIndex {
		slice[k] = leftArray[i]
		i++
		k++
	}

	for j < rightIndex {
		slice[k] = rightArray[j]
		j++
		k++
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
