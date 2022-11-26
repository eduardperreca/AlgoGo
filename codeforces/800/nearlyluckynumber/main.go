package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var count int
	for n != 0 {
		if n%10 == 4 || n%10 == 7 {
			count++
		}
		n /= 10
	}
	if count == 4 || count == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
