package main

import "fmt"

func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
			fmt.Println(array, i, j)
		}
	}
	return array
}

func main() {

	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(selectionSort(array))
}
