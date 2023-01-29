package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Arch struct {
	destination int
	brightness  int
}

type Graph struct {
	n         int
	m         int
	adjacents [][]Arch
}

func shortestPath(g Graph, from int, to int) int {

	if g.adjacents[from] == nil {
		return -1
	}

	count := 0
	visited := make(map[int]bool)
	v := from
	minTo := -1
	min := 9999
	for v != to {
		visited[v] = true
		for _, k := range g.adjacents[v] {
			if k.brightness < min {
				min = k.brightness
				minTo = k.destination
			}
		}

		if !visited[minTo] {
			count++
			v = minTo
		} else {
			break
		}
	}

	if v == to {
		fmt.Println("ci sono volute: ", count, "gallerie")
		return count
	} else {
		fmt.Println(-1)
		return -1
	}
}

func main() {
	var g Graph

	n, m, h, s := 0, 0, 0, 0
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	i := 0
	appendSlice := []int{}
	// lettura input tostarella
	for b.Scan() {
		if i == 0 {
			slice := strings.Split(b.Text(), " ")
			for _, k := range slice {
				numb, _ := strconv.Atoi(k)
				appendSlice = append(appendSlice, numb)
			}
			n, m, h, s = appendSlice[0], appendSlice[1], appendSlice[2], appendSlice[3]
			// setup adjacent list
			g.adjacents = make([][]Arch, n+1)
			g.n, g.m = n, m
			fmt.Println(n, m, h, s) // h harmony start s sarah end
		} else {
			fmt.Println(b.Text())
			slice := strings.Split(b.Text(), " ")
			appendSlice = make([]int, 0)
			for _, k := range slice {
				numb, _ := strconv.Atoi(k)
				appendSlice = append(appendSlice, numb)
			}
			from, to, brightness := appendSlice[0], appendSlice[1], appendSlice[2]
			g.adjacents[from] = append(g.adjacents[from], Arch{to, brightness})
			g.adjacents[to] = append(g.adjacents[to], Arch{from, brightness})
		}
		i++
	}

	// for i := 1; i < g.n+1; i++ {
	// 	fmt.Println(i, ":", g.adjacents[i])
	// }

	t := shortestPath(g, h, s)
	fmt.Println(t)
}
