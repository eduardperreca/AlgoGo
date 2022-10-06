package main

import "fmt"

func main() {
	var check int
	slice := []int{4, 5, 1, 3, 6, 2}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			check++
		}
	}
	fmt.Println(check)
}
