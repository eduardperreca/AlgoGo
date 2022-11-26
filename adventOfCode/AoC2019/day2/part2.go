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
		// fmt.Println(scanner.Text())
		needed := strings.Split(scanner.Text(), ",")
		for _, k := range needed {
			numb, _ := strconv.Atoi(k)
			sliceIntegers = append(sliceIntegers, numb)
		}
	}

	for j := 0; j < 100; j++ {
		found := false
		for k := 0; k < 100; k++ {
			running := make([]int, len(sliceIntegers))
			copy(running, sliceIntegers)
			running[1] = j
			running[2] = k

			for i := 0; i < len(running); i += 4 {
				if running[i] == 99 {
					break
				}
				if running[i] == 1 {
					running[running[i+3]] = running[running[i+1]] + running[running[i+2]]
				}
				if running[i] == 2 {
					running[running[i+3]] = running[running[i+1]] * running[running[i+2]]
				}
			}
			if running[0] == 19690720 {
				fmt.Println("ciao")
				fmt.Println(100*j + k)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}
