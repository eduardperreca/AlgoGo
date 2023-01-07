package main

import "fmt"

func main() {
	n := 10
	a, b := 1, 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	fmt.Println(b)
}
