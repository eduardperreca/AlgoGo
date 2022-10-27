package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Previous *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}

func addNewNode(list *List, value int) {
	node := newNode(value)
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		list.Tail.Previous = node
	}
}

func main() {

	var l List
	fmt.Println(l)
	addNewNode(&l, 1)
	addNewNode(&l, 2)
	addNewNode(&l, 3)
	addNewNode(&l, 4)
	fmt.Println(l)
	fmt.Println(l.Head, " primo")
	fmt.Println(l.Head.Next, " secondo")
	fmt.Println(l.Tail, " ultimo")
	fmt.Println(l.Tail.Next, " nil")
}
