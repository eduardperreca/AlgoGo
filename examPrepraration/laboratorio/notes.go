// fibonacci con dp, salvo gli stati in un array
package main

func main(){

	var dp = make([]int, 10)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < 10; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	println(dp[9])
	
}

// time complexity: O(n)
// space complexity: O(n)


// --------------------------------------------- //


// knapsack problem with dp (dynamic programming) 
// time complexity: O(n*W)
// space complexity: O(n*W)

package main

import "fmt"

func knap(wt, val []int, W int) int {
	n := len(wt)
	dp := make([][]int, n+1)

	// creo una matrice di n+1 righe e W+1 colonne	
	for i := range dp {
		dp[i] = make([]int, W+1)
	}


	// dp[i][j] = max value of the knapsack with weight j and the first i items
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
	// wt = weight
	wt := []int{1, 3, 4, 5}
	// val = value
	val := []int{1, 4, 5, 7}
	// W = max weight		
	W := 12
	fmt.Println(knap(wt, val, W))
}


// --------------------------------------------- //


// shortest path in a matrix with dp (dynamic programming)
// time complexity: O(n*m)
// space complexity: O(n*m)

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


// --------------------------------------------- //


// longest common subsequence with dp (dynamic programming)
// time complexity: O(n*m)
// space complexity: O(n*m)

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	
	var a = "AGGTAB"
	var b = "GXTXAYB"
	var n = len(a)
	var m = len(b)

	var dp = make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(dp[n][m])

}

// --------------------------------------------- //

// longest increasing subsequence with dp (dynamic programming)
// time complexity: O(n^2)
// space complexity: O(n)

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	
	var a = []int{10, 22, 9, 33, 21, 50, 41, 60}
	var n = len(a)

	var dp = make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(dp[n-1])

}

// --------------------------------------------- //

// longest palindromic subsequence with dp (dynamic programming)
// time complexity: O(n^2)
// space complexity: O(n^2)

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	
	var a = "GEEKSFORGEEKS"
	var n = len(a)

	var dp = make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			if a[i] == a[j] && l == 2 {
				dp[i][j] = 2
			} else if a[i] == a[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(dp[0][n-1])

}

// --------------------------------------------- //
