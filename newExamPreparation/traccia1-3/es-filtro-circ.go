package main

import "fmt"

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

type linkedList struct {
	Head *Node
	Tail *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *linkedList, k int) {
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

func stampaDaZero(p *Node) {
	visited := make(map[int]bool)
	checkZero := false
	for {
		if p.Val == 0 {
			checkZero = true
		}
		if checkZero {
			if !visited[p.Val]{
				visited[p.Val] = true
				fmt.Print(p.Val, " ")
			} else {
				return
			}
		}
		p = p.Next
	}

}

func main() {
	l := &linkedList{}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		val := 0
		fmt.Scan(&val)
		addNewNode(l, val)
	}
	stampaDaZero(l.Head)
}
