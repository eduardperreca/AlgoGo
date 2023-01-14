package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	matrix := [][]int{
		{7, 1, 8, 4, 4},
		{9, 1, 8, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 1},
	}

	dp := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	dp[0][0] = matrix[0][0]

	for i := 1; i < len(matrix); i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}

	for j := 1; j < len(matrix[0]); j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}

	println(dp[len(matrix)-1][len(matrix[0])-1] - matrix[0][0])

}
