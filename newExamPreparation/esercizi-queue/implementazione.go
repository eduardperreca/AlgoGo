package main

import "fmt"

type Node struct {
	Previous *Node
	Val      int
}

type Queue struct {
	Tail *Node
	Head *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *Queue, k int) {
	node := newNode(k)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Previous = l.Tail
		l.Tail = node
	}
}

func queue(l *Queue, k int) {
	addNewNode(l, k)
}

func dequeue(l *Queue) int {

	node := l.Tail
	val := l.Head.Val
	head := l.Head
	for node.Previous != head {
		if node.Previous.Previous == head {
			l.Head = node
			fmt.Println(node.Val, head.Val)
		}
		node = node.Previous
	}

	return val
}

func main() {

	l := &Queue{}
	queue(l, 12)
	queue(l, 23)
	queue(l, 1)
	queue(l, 10)

	node := l.Tail
	for node != nil{
		fmt.Println(node.Val)
		node = node.Previous
	}


	fmt.Println(dequeue(l))

	node2 := l.Tail
	for node2 != nil{
		fmt.Println(node2.Val)
		node2 = node2.Previous
	}
	
}
