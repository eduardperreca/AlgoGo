package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, _ := os.Open("nomi.txt")

	b := bufio.NewScanner(f)
	i := 0
	for b.Scan() {
		i++
	}

	fmt.Println(i)

}
