package main

import "fmt"

func selectionSortRec(array []int) []int {
	if len(array) == 1 {
		return array
	}
	min := array[0]
	minIndex := 0
	for i, v := range array {
		if v < min {
			min = v
			minIndex = i
		}
	}
	array[0], array[minIndex] = array[minIndex], array[0]
	fmt.Println(array)
	return append([]int{min}, selectionSortRec(array[1:])...)
}

func main() {

	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(selectionSortRec(array))
}
