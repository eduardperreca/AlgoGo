package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		numb, _ := strconv.Atoi(scanner.Text())
		partial := (numb / 3) - 2
		fmt.Println(partial)
		for partial/3-2 > 0 {
			partial = partial/3 - 2
			numb += partial
			partial = partial/3 - 2
		}
		total += numb
	}
	fmt.Println(total)
}
