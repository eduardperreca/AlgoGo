package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	fmt.Println(left, right)

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	middle := len(array) / 2
	left := mergeSort(array[:middle])
	right := mergeSort(array[middle:])
	return merge(left, right)
}

func main() {
	
	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(mergeSort(array))

}
