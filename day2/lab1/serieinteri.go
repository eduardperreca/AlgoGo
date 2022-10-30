package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}

	fmt.Println(stranoProdotto(numeri))
}

func stranoProdotto(numeri []int) int {
	prod := 1
	for _, n := range numeri{
		if n > 7 && n % 4 == 0{
			prod *= n
		}
	}
	return prod
}