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
	fmt.Println(dp)
	return dp[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}
	W := 12
	fmt.Println(knap(wt, val, W))
}
