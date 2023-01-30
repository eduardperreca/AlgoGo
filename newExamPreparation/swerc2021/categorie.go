package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}

	sort.Ints(d)

	count := 0
	for k := 0; k < n; k++ {
		numGreaterEqual := 0
		numLess := 0
		for i := 0; i < n; i++ {
			if d[i] >= d[k] {
				numGreaterEqual++
			} else {
				numLess++
			}
		}
		if numGreaterEqual == numLess {
			count++
		}
	}

	fmt.Println(count)
}
