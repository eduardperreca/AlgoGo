package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	var stringa string
	fmt.Scan(&stringa)

	var counter int

	for i := 0; i < len(stringa)-1; i++ {
		if stringa[i] == stringa[i+1] {
			counter++
		}
	}
	fmt.Println(counter)

}
