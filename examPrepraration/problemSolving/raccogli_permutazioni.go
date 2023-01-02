package main

import "fmt"

func main() {

	permutazione := []int{4, 5, 1, 3, 6, 2}
	var count = 0
	occ := make(map[int]bool)
	for _, el := range permutazione {
		occ[el] = true
		fmt.Println(occ[el+1])
		if occ[el+1] {
			count++
		}
	}
	fmt.Println(count)
}
