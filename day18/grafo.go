package main

import "fmt"

type grafo struct {
	n         int
	adiacenti []*LinkedList
}

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func newGraph(n int) *grafo {
	g := new(grafo)
	g.n = n
	g.adiacenti = make([]*LinkedList, n)
	for i := 0; i < n; i++ {
		g.adiacenti[i] = new(LinkedList)
	}
	return g
}

func (l *LinkedList) add(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func readGraph() *grafo {
	var n int
	fmt.Scan(&n)
	g := newGraph(n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		g.adiacenti[a].add(b)
		g.adiacenti[b].add(a)
	}
	return g
}

func printGraph(g *grafo) {
	for i := 0; i < g.n; i++ {
		fmt.Printf("%d: ", i)
		current := g.adiacenti[i].head
		for current != nil {
			fmt.Printf("%d ", current.value)
			current = current.next
		}
		fmt.Println()
	}
}

func checkEdge(g *grafo, a, b int) bool {
	current := g.adiacenti[a].head
	for current != nil {
		if current.value == b {
			return true
		}
		current = current.next
	}
	return false
}

func main() {

}
