package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	for b.Scan() {

	}

}
