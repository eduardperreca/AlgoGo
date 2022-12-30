package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDivisible(slice []int) int {
	fmt.Println(slice, "slice")
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i]%slice[j] == 0 {
				return slice[i] / slice[j]
			}
			if slice[j]%slice[i] == 0 {
				return slice[j] / slice[i]
			}
		}
	}
	return 0
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var sum int
	var slice []int
	for b.Scan() {
		slice = []int{}
		splitted := strings.Split(b.Text(), "\t")

		fmt.Println(splitted, "splitted")
		for _, k := range splitted {
			fmt.Println(k, "k")
			numb, _ := strconv.Atoi(k)
			slice = append(slice, numb)
		}

		sum += findDivisible(slice)

	}
	fmt.Println(sum)
}
