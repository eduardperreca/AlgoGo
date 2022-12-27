package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	countR := 0
	countL := 0
	for b.Scan() {
		splitted := strings.Split(b.Text(), ",")
		for _, s := range splitted {
			if s[0] == 'R' {
				numb, _ := strconv.Atoi(s[1:])
				countR += numb
			} else {
				numb, _ := strconv.Atoi(s[1:])
				countL += numb
			}
		}
	}
	fmt.Println(abs(countR) + abs(countL))
}
