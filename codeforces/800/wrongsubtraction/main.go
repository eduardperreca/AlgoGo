package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 0; i < b; i++ {
		if a%10 == 0 {
			a /= 10
		} else {
			a--
		}
	}
	fmt.Println(a)
}
