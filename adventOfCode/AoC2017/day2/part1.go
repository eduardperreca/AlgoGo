package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

func max(slice []int) int {
	max := 0
	for _, k := range slice {
		if k > max {
			max = k
		}
	}
	return max
}

func min(slice []int) int {
	min := 100000000
	for _, k := range slice {
		if k < min {
			min = k
		}
	}
	return min
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var sum int
	var slice []int
	for b.Scan() {
		slice = []int{}
		splitted := strings.Split(b.Text(), "\t")
		
		fmt.Println(splitted)
		for _, k := range splitted {
			slice = append(slice, atoi(k))
		}

		sum += max(slice) - min(slice)
	}
	fmt.Println(sum)
}
