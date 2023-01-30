package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	count := 0
	for b.Scan() {
		var min, max int
		var rule string
		fmt.Println(b.Text())
		slice := strings.Split(b.Text(), " ")
		numbs := slice[0]
		numbsArr := strings.Split(numbs, "-")
		min, _ = strconv.Atoi(numbsArr[0])
		max, _ = strconv.Atoi(numbsArr[1])
		rules := slice[1]
		rule = string(rules[0])
		parole := make([]string, 0)
		for _, k := range slice[2] {
			parole = append(parole, string(k))
		}
		if parole[min-1] == rule && parole[max-1] == rule {
			count++
		}

		// fmt.Println(rule, min, max)
	}
	fmt.Println(count)
}
