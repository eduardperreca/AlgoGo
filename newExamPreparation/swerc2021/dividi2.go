package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	ans := int(1e9)
	for k := 1; k < n; k++ {
		sum, sum2 := 0, 0

		for i := 0; i < k; i++ {
			sum += arr[i]
		}
		for j := k; j < n; j++ {
			sum2 += arr[j]
		}
		diff := sum - sum2
		fmt.Println(diff)
		diff = abs(diff)
		ans = min(ans, diff)

	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
