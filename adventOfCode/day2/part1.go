package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	aim := 0
	counterVertical := 0
	counterHorizontal := 0
	for b.Scan() {

		dummy := strings.Split(b.Text(), " ")
		text := dummy[0]
		numb, _ := strconv.Atoi(dummy[1])
		println(text, numb)
		if text == "forward" {
			counterHorizontal += numb
			counterVertical += numb * aim
		} else if text == "down" {
			aim += numb
		} else {
			aim -= numb
		}
		println(counterHorizontal, counterVertical, aim)
	}
	println(counterHorizontal * counterVertical)
}
