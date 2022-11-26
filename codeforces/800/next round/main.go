package main

import (
	"fmt"
)

func main() {
	var n, k, a int
	fmt.Scan(&n, &k)
	p := 0
	x := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		if i == k {
			p = a
		}

		if (a >= p) && (a > 0) {
			x += 1
		}
	}
	fmt.Print(x)
}
