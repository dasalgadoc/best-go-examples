package main

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := linearSearch(slice, 3)

	println(index)
}

func linearSearch(slice []int, target int) int {
	for index, item := range slice {
		if item == target {
			return index
		}
	}

	return -1
}
