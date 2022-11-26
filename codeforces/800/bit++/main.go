package main

import (
	"bufio"
	"os"
)

func main() {

	b := bufio.NewScanner(os.Stdin)
	tot := 0
	damn := 0
	for b.Scan() {
		for _, c := range b.Text() {
			damn += int(c) - 48
			
		}
	}

}
