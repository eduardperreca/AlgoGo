package main

import "fmt"

func knap(wt, val []int, W int) int {
	n := len(wt)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {
			if wt[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// se ho spazio per il peso, prendo il massimo tra il valore
				dp[i][j] = max(val[i-1]+dp[i-1][j-wt[i-1]], dp[i-1][j])
			}
		}
	}
	for _, k := range dp {
		fmt.Println(k)
	}
	return dp[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	wt := []int{20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40}
	val := []int{100, 107, 124, 133, 139, 155, 172, 178, 184, 190, 195}
	W := 140
	fmt.Println(knap(wt, val, W))
}
