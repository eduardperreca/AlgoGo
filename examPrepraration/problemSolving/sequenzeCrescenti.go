package main

import "fmt"

func main() {
	counter := 0
	sliceInter := []int{9, 1, 3, 5, 2, 0, 8, 6}
	for i := 1; i < len(sliceInter)-1; i++ {
		if sliceInter[i] < sliceInter[i-1] && sliceInter[i] < sliceInter[i+1] {
			counter++
		}
	}
	fmt.Println(counter)
}
