package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	counter := 0
	slice := make([]int, 0)
	for b.Scan() {
		numb, _ := strconv.Atoi(b.Text())
		slice = append(slice, numb)
	}
	for i := 1; i<len(slice); i++{
		if slice[i] > slice[i-1]{
			counter++
		}
	}
	println(counter)
}
