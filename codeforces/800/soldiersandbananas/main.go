package main

import "fmt"

func main() {

	var k, n, w int

	fmt.Scan(&k, &n, &w)
	
	var sum int

	for i := 1; i <= w; i++ {
		sum += i * k
	}
	if sum > n {
		fmt.Println(sum - n)
	} else {
		fmt.Println(0)
	}

}
