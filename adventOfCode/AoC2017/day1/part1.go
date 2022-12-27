package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	var slice []int
	for b.Scan() {
		for _, k := range b.Text() {
			x, _ := strconv.Atoi(string(k))
			slice = append(slice, x)
		}
	}

	var sum int
	for i := 0; i < len(slice); i++ {
		if slice[i] == slice[(i+1)%len(slice)] {
			sum += slice[i]
		}
	}
	fmt.Println(sum)
}
