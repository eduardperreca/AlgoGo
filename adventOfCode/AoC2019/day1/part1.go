package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		numb, _ := strconv.Atoi(scanner.Text())
		numb = (numb / 3) - 2
		total += numb
	}
	print(total)
}
