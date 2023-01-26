package main

type Node struct {
	Val  string
	Next *Node
}

type Stack struct {
	Head *Node
}

func newNode(k string) *Node {
	return &Node{Val: k}
}

func addNewNode(l *Stack, k string) {
	node := newNode(k)
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func main() {
	var 
}
