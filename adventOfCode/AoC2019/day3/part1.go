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
	scanner := bufio.NewScanner(f)
	needed := []string{}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		needed = strings.Split(scanner.Text(), ",")
	}

	tot := 0
	for _, n := range needed {
		if n[0] == 'R' {
			fmt.Println("Right")
			num, _ := strconv.Atoi(n[1:])
			tot += num
		} else if n[0] == 'L' {
			fmt.Println("Left")
			num, _ := strconv.Atoi(n[1:])
			tot -= num
		} else if n[0] == 'U' {
			fmt.Println("Up")
			num, _ := strconv.Atoi(n[1:])
			tot += num
		} else if n[0] == 'D' {
			fmt.Println("Down")
			num, _ := strconv.Atoi(n[1:])
			tot -= num
		}
	}
	fmt.Println(tot)

}
