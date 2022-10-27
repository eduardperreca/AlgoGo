package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}

func addNewNodeLast(l *List, value int) {
	node := newNode(value)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		node.Previous = l.Tail
		l.Tail = node
	}
}

func addNewNode(list *List, value int) {
	node := newNode(value)
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.Previous = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}
}

func removeAtStart(l *List) {
	l.Head = l.Head.Next
}

func removeAtEnd(l *List) {
	l.Tail = l.Tail.Previous
}

func removeAtPosition(l *List, position int) {
	var i int
	for node := l.Head; node != nil; node = node.Next {
		if i == position {
			node.Previous.Next = node.Next
			node.Next.Previous = node.Previous
		}
		i++
	}
}

func printBackwards(l *List) {
	for node := l.Tail; node != nil; node = node.Previous {
		fmt.Println(node.Value)
	}
}

func contains(l *List, value int) int {
	for node := l.Head; node != nil; node = node.Next {
		if node.Value == value {
			return 1
		}
	}
	return 0
}

func printList(l *List) {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

func main() {

	var l List

	addNewNode(&l, 1)
	addNewNode(&l, 2)
	addNewNode(&l, 3)
	addNewNode(&l, 4)
	addNewNode(&l, 5)
	addNewNode(&l, 6)
	addNewNode(&l, 7)
	addNewNode(&l, 8)
	addNewNode(&l, 9)
	addNewNode(&l, 10)

	printList(&l)

	fmt.Println(l.Tail.Next)
	
	// fmt.Println(l.Head.Value, "Testa")
	// fmt.Println(l.Head.Next.Value, "Next")
	// fmt.Println(l.Tail.Value, "Coda")
	// fmt.Println(l.Tail.Previous.Value, "Previous dell'ultimo")


}
