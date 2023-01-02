package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	insertionSort(arr)
	fmt.Println(arr)
}
