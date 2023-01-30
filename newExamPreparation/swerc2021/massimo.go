package main

import "fmt"

func max(slice []int) int {
	max := 0
	for _, k := range slice {
		if k > max {
			max = k
		}
	}
	return max
}

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	m := max(arr)

	for i, k := range arr {
		if k == m {
			fmt.Println(max(append(arr[:i], arr[i+1:]...)))
		} else {
			fmt.Println(m)
		}
	}

}
