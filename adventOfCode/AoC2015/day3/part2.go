package main

import (
	"bufio"
	"fmt"
	"os"
)

type house struct {
	x, y int
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	visited := make(map[house]bool)
	visited[house{0, 0}] = true
	visited2 := make(map[house]bool)
	visited2[house{0, 0}] = true
	x, y := 0, 0
	x2, y2 := 0, 0
	for b.Scan() {

		for i, c := range b.Text() {
			if i%2 == 0 && i != 0 {
				switch c {
				case '<':
					x--
				case '>':
					x++
				case '^':
					y++
				case 'v':
					y--
				}
				visited[house{x, y}] = true
			} else {
				fmt.Println("else")
				switch c {
				case '<':
					x2--
				case '>':
					x2++
				case '^':
					y2++
				case 'v':
					y2--
				}
				visited2[house{x2, y2}] = true
			}
			fmt.Println(i)

		}
	}

	fmt.Println(len(visited), len(visited2), len(visited)+len(visited2))
}
