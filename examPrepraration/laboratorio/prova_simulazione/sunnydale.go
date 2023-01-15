package main

import "fmt"

type arcoUscente struct {
	v int // nodo di destinazione
	l int // luminosità dell'arco
}

type Graph struct {
	nodes int
	arcs  int
	adj   [][]arcoUscente
}

func camminoMinimo(g Graph, from int, to int) int {
	if g.adj == nil {
		return -1
	}

	vertex := from // nodo corrente
	visited := make(map[int]bool)
	count := 0

	for vertex != to {
		visited[vertex] = true
		// trovo il nodo con luminosità minima
		min := 1000000
		minTo := -1
		for i := 0; i < len(g.adj[vertex]); i++ {
			if !visited[g.adj[vertex][i].v] && g.adj[vertex][i].l < min {
				min = g.adj[vertex][i].l
				minTo = g.adj[vertex][i].v
			}
		}

		fmt.Println("min", min, "minTo", minTo)

		if visited[minTo] {
			break
		} else {
			vertex = minTo
			count++
		}
	}

	if vertex == to {
		fmt.Println("count numero gallerie passate: ", count)
		return count
	} else {
		return -1
	}
}

func main() {

	var n int // numero di nodi
	var m int // numero di gallerie
	var h int // casa della tipa -> nodo di partenza
	var s int // casa di Sarah -> nodo di arrivo

	fmt.Scan(&n, &m, &h, &s)
	fmt.Println(n, m, h, s)

	// inizializzo il grafo
	var g Graph
	g.nodes = n
	g.arcs = m
	g.adj = make([][]arcoUscente, n+1)

	// leggo gli archi
	for i := 0; i < m; i++ {
		var a, b, l int
		fmt.Scan(&a, &b, &l)
		fmt.Println(a, b, l)
		g.adj[a] = append(g.adj[a], arcoUscente{b, l})
		g.adj[b] = append(g.adj[b], arcoUscente{a, l})
	}

	// stampo il grafo
	for i := 1; i <= n; i++ {
		fmt.Println("nodo", i)
		for j := 0; j < len(g.adj[i]); j++ {
			fmt.Println("  ", g.adj[i][j].v, g.adj[i][j].l)
		}
	}

	fmt.Println(camminoMinimo(g, h, s))

}
