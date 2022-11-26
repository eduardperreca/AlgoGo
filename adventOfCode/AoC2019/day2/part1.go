package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sliceIntegers := []int{}

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		needed := strings.Split(scanner.Text(), ",")
		for _, k := range needed {
			numb, _ := strconv.Atoi(k)
			sliceIntegers = append(sliceIntegers, numb)
		}
	}


	for i := 0; i < len(sliceIntegers); i += 4 {
		if sliceIntegers[i] == 99 {
			break
		}
		if sliceIntegers[i] == 1 {
			sliceIntegers[sliceIntegers[i+3]] = sliceIntegers[sliceIntegers[i+1]] + sliceIntegers[sliceIntegers[i+2]]
		}
		if sliceIntegers[i] == 2 {
			sliceIntegers[sliceIntegers[i+3]] = sliceIntegers[sliceIntegers[i+1]] * sliceIntegers[sliceIntegers[i+2]]
		}
	}
	fmt.Println(sliceIntegers[0])
}
