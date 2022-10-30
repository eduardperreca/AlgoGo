package main

import "fmt"

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	left := make([]int, 0, len(array))
	right := make([]int, 0, len(array))

	for _, v := range array[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main(){

	arr := []int{1, 12, 4, 8, 6, 32, 9}
	fmt.Println(arr)
	t := quickSort(arr)
	fmt.Println(t)
}