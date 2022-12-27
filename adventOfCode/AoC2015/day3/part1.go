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

	x, y := 0, 0

	for b.Scan() {

		for _, c := range b.Text() {

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

		}

	}

	fmt.Println(len(visited))

}
