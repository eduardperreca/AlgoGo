package main

import "fmt"

type Node struct {
	Next     *Node
	Previous *Node
	Val      int
}

type circList struct {
	Head *Node
	Tail *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *circList, k int) {
	node := newNode(k)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Previous = node
		l.Head = node
		l.Head.Previous = l.Tail
		l.Tail.Next = l.Head
	}
}

func stampaDaZero(node *Node) {

	visited := make(map[int]bool)

	reached := false
	for {

		if node.Val == 0 {
			reached = true
		}
		if reached {
			if !visited[node.Val] {
				visited[node.Val] = true
				fmt.Print(node.Val, " ")
			} else {
				return
			}
		}
		node = node.Next
	}
}

func main() {
	l := &circList{}
	addNewNode(l, 7)
	addNewNode(l, 1)
	addNewNode(l, 0)
	addNewNode(l, -2)
	addNewNode(l, 3)

	node := l.Head
	stampaDaZero(node)

}
