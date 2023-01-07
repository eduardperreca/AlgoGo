package main

import "fmt"

func min(x, y, a, b int) int {
	if x < y && x < a && x < b {
		return x
	} else if y < x && y < a && y < b {
		return y
	} else if a < x && a < y && a < b {
		return a
	} else {
		return b
	}
}

func main() {

	caverna := [][]int{
		{7, 1, 8, 4, 4},
		{9, 1, 8, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 1},
	}

	for i := 0; i < len(caverna)-1; i++ {
		for j := 0; j < len(caverna[i])-1; j++ {

			fmt.Println(caverna[i+1][j])

		}
	}

}
