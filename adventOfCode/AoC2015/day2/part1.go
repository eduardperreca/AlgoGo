package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func getRes(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	res := 0
	for b.Scan() {

		splitted := strings.Split(b.Text(), "x")

		l, _ := strconv.Atoi(splitted[0])
		w, _ := strconv.Atoi(splitted[1])
		h, _ := strconv.Atoi(splitted[2])

		res += getRes(l, w, h)

	}

	fmt.Println(res)
}
