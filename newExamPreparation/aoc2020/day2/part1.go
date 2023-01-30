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
		var mappa = make(map[string]int)
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
		for _, k := range slice[2] {
			mappa[string(k)]++
		}
		for k, v := range mappa {
			if k == rule {
				if v >= min && v <= max {
					fmt.Println("giusto")
					count++
				}
			}
		}
		fmt.Println(rule, min, max, mappa)
	}
	fmt.Println(count)
}
