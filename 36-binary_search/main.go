package main

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := binarySearch(slice, 3)

	println(index)
}

func binarySearch(slice []int, target int) int {
	first := 0
	last := len(slice) - 1

	for first <= last {
		mid := (first + last) / 2

		if slice[mid] < target {
			first = mid + 1
		} else if slice[mid] > target {
			last = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
