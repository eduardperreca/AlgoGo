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

	for i := 0; i < node; i++ {
		fmt.Println(i, g.adjacents[i])
	}

	matrix := pesiToMatrice(g)
}

func pesiToMatrice(g Graph) [][]int {

	matrix := [][]int{}
	for i := 0; i<len(g.adjacents); i++{
		
	}

}
