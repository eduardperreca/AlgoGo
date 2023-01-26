package main

import "fmt"

type linkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Next     *Node
	Previous *Node
	Val      int
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

func main() {

	l := &linkedList{}
	addNewNode(l, 7)
	addNewNode(l, 1)
	addNewNode(l, 0)
	addNewNode(l, -2)
	addNewNode(l, 3)

	node := l.Head
	for {
		fmt.Println(node.Val)
		node = node.Previous
	}
}
