package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	numb, _ := strconv.Atoi(b.Text())

	if numb%2 == 0 && numb != 2 {
		println("YES")
	} else {
		println("NO")
	}

}
