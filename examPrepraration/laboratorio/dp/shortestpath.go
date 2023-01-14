package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	var matrix = [][]int{
		{7, 1, 8, 4, 4},
		{9, 1, 8, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 1},
	}

	var n = len(matrix)
	var m = len(matrix[0])

	var dp = make([][]int, n)

	// inizializzo la matrice a 0 per n volte quante le righe
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// assegno il valore iniziale come il valore della matrice di partenza
	dp[0][0] = matrix[0][0]
	fmt.Println(dp[0][0], matrix[0][0])

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
		// fmt.Println(dp[i][0], matrix[i][0], dp[i-1][0])
	}
	// fmt.Println(dp)

	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
		// fmt.Println(dp[0][j], matrix[0][j], dp[0][j-1])
	}
	// fmt.Println(dp)

	// finisco di riempire la matrice con i casi calcolati precedentemente prendendo il minore
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fmt.Println(dp[0][0], dp[1][1])
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
			// fmt.Println(dp[i][j], matrix[i][j], dp[i-1][j], dp[i][j-1)
		}
	}
	fmt.Println(dp)

	// fmt.Println(matrix)
	// fmt.Println(dp)
	fmt.Println(dp[n-1][m-1] - matrix[0][0])

}
