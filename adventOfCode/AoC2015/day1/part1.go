package main

import (
	"bufio"
	"fmt"
	"os"
)

// ( -> +1
// ) -> -1

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	var mappa = make(map[string]int)
	for b.Scan() {

		for i, k := range b.Text() {
			mappa[string(k)]++
			if (mappa["("]-mappa[")"]) == -1 {
				fmt.Println(i + 1)
				break
			}

		}

	}
	// fmt.Println(mappa["("] - mappa[")"])
}
