package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	slice := make([]int, 0)
	for b.Scan() {
		numb, _ := strconv.Atoi(b.Text())
		slice = append(slice, numb)
	}

	counter := 0
	for i := 0; i < len(slice)-3; i++ {

		sum := slice[i] + slice[i+1] + slice[i+2]
		succ := slice[i+1] + slice[i+2] + slice[i+3]

		if sum < succ {
			counter++
		}
	}
	println(counter)
}
