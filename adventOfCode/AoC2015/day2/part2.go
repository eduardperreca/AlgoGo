package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}
	return c
}

func getRes(l, w, h int) int {
	return 2*l + 2*w + 2*h - 2*max(l, w, h) + l*w*h
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
