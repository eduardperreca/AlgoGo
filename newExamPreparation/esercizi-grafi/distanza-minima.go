package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Arch struct {
	direction int
	weight    int
}

type Graph struct {
	nodes     int
	bows      int
	adjacents [][]Arch
}

type Node struct {
	to     int
	weight int
}

func shortestPath(g Graph, from int, to int) int {
	visited := make(map[int]bool)
	queue := []Arch{}

	queue = append(queue, Arch{0, 0})
	visited[queue[0].direction] = true
	i := 0
	sum := 0
	for len(queue) > 0 {
		fmt.Println(queue)
		node := queue[0]
		fmt.Println("visiting: ", node)
		min := 999
		queue = queue[1:]
		fmt.Println(queue)
		nodeMin := Arch{0, 0}
		if node.direction == to {
			fmt.Println("end reached we spent: ", sum, "steps wow")
			return sum
		}
		for i := range g.adjacents[node.direction] {
			fmt.Println(g.adjacents[node.direction][i])
			if g.adjacents[node.direction][i].weight < min && !visited[g.adjacents[node.direction][i].direction] {
				visited[g.adjacents[node.direction][i].direction] = true
				min = g.adjacents[node.direction][i].weight
				nodeMin = g.adjacents[node.direction][i]
			}
		}
		fmt.Println(min, "min", nodeMin)
		queue = append(queue, nodeMin)
		sum += nodeMin.weight
		fmt.Println(queue)
		i++
	}

	return -1

}

func main() {

	var g Graph
	node, bow := 0, 0

	f, _ := os.Open("input2.txt")
	b := bufio.NewScanner(f)
	i := 0
	dummySlice := []int{}
	for b.Scan() {
		if i == 0 {
			split := strings.Split(b.Text(), " ")
			for _, k := range split {
				numb, _ := strconv.Atoi(k)
				dummySlice = append(dummySlice, numb)
			}
			node, bow = dummySlice[0], dummySlice[1]
			g.nodes = node
			g.bows = bow
			g.adjacents = make([][]Arch, node+1)
		} else {
			dummySlice = make([]int, 0)
			split := strings.Split(b.Text(), " ")
			for _, k := range split {
				numb, _ := strconv.Atoi(k)
				dummySlice = append(dummySlice, numb)
			}
			from, to, weight := dummySlice[0], dummySlice[1], dummySlice[2]
			g.adjacents[from] = append(g.adjacents[from], Arch{to, weight})
			g.adjacents[to] = append(g.adjacents[to], Arch{from, weight})
		}
		i++
	}

	shortestPath(g, 0, node-1)

	// for i := 0; i < node; i++ {
	// 	fmt.Println(i, g.adjacents[i])
	// }

}
