package main

import "fmt"

func insertionSortRec(array []int) []int {
	if len(array) == 1 {
		return array
	}

	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return append(insertionSortRec(array[:len(array)-1]), array[len(array)-1])
}

func main() {
	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(insertionSortRec(array))
}
