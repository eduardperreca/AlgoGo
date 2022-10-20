package main

import "fmt"

func main() {

	arr := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
