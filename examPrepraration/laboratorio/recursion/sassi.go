package main

import "fmt"

func sassi(h int) int {
	if h == 1 {
		return 1
	}
	return h*h+sassi(h-1) 
}

func main() {

	h := 0
	fmt.Scan(&h)
	fmt.Println("i sassi sono per", h, "sono: ", sassi(h))
}
