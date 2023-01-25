package main

import "fmt"

type Node struct {
	Next     *Node
	Previous *Node
	Val      int
}

type linkedList struct {
	Tail *Node
	Head *Node
}

func stampaDaZero(p *Node) {
	toCheck := false
	node := p
	var visited = make(map[int]bool)
	for {
		if node.Val == 0 {
			toCheck = true

		}
		if toCheck {
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

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNodeHead(p *linkedList, k int) {
	node := newNode(k)
	if p.Head == nil {
		p.Head = node
		p.Tail = node
	} else {
		node.Next = p.Head
		p.Head.Previous = node
		p.Head = node
		p.Head.Previous = p.Tail
		p.Tail.Next = p.Head
	}
}

func main() {

	p := &linkedList{}

	addNewNodeHead(p, 7)
	addNewNodeHead(p, 1)
	addNewNodeHead(p, 0)
	addNewNodeHead(p, -2)
	addNewNodeHead(p, 3)

	node := p.Head
	stampaDaZero(node)
}
