package main

import (
	"fmt"
	"math"
)

func main() {
	// Lettura della lunghezza dell'array
	var n int
	fmt.Scanf("%d", &n)

	// Lettura dell'array
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	// Calcolo della differenza minima
	ans := math.MaxInt32
	for k := 1; k < n; k++ {
		sum1 := 0
		sum2 := 0
		for i := 0; i < k; i++ {
			sum1 += a[i]
		}
		for i := k; i < n; i++ {
			sum2 += a[i]
		}
		diff := int(math.Abs(float64(sum1 - sum2)))
		ans = int(math.Min(float64(ans), float64(diff)))
	}

	// Stampa della risposta
	fmt.Println(ans)
}
