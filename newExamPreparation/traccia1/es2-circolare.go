package main

import (
	"fmt"
)

type Node struct {
	Next     *Node
	Previous *Node
	Val      int
}

type linkedList struct {
	Tail *Node
	Head *Node
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

func f2(p *Node, k int) int {

	start := p
	end := start.Next
	a := 1
	for start.Val != end.Val {
		a++
		if a == k {
			fmt.Println(start.Val)
		}
		start = start.Previous
	}

	return a
}

// lineare

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func sizeOf(p *Node) int {
	var visited = make(map[int]bool)
	for {
		if !visited[p.Val] {
			visited[p.Val] = true
		} else {
			break
		}
		p = p.Next
	}
	return len(visited)
}

func sposta(p *Node) {
	node := p
	var visited = make(map[int]bool)
	for {
		if len(visited) == sizeOf(node) {
			return
		}
		if !visited[node.Val] {
			visited[node.Val] = true
			n := abs(node.Val)
			for i := 0; i < n; i++ {
				if node.Val > 0 {
					node.Val, node.Previous.Val = node.Previous.Val, node.Val
				} else if node.Val < 0 {
					node.Val, node.Next.Val = node.Next.Val, node.Val
				}
			}
		}
		node = node.Next
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
	sposta(node)
	for i := 0; i < 5; i++ {
		fmt.Println(node.Val)
		node = node.Next
	}
}
