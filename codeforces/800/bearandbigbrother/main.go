package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a, &b)

	var counter int
	for a <= b {
		a *= 3
		b *= 2
		counter++
	}
	fmt.Println(counter)
}
