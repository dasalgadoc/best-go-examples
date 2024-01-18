package main

import "sort"

func isSliceSortedStd(slice []int) bool {
	return sort.IsSorted(sort.IntSlice(slice))
}

func isSliceSortedVanilla(slice []int) bool {
	length := len(slice)
	if length <= 1 {
		return true
	}

	return slice[0] <= slice[1] && isSliceSortedVanilla(slice[1:])
}
