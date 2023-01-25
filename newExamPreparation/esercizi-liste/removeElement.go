package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type linkedList struct {
	Head *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *linkedList, k int) {
	node := newNode(k)
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func removeElement(node *Node, k int) {

	for node.Next != nil {
		if node.Next.Val == k {
			node.Next = node.Next.Next
		}
		node = node.Next
	}

}

func main() {

	l := &linkedList{}
	addNewNode(l, 5)
	addNewNode(l, 10)
	addNewNode(l, 12)
	addNewNode(l, 4)
	addNewNode(l, 17)

	node := l.Head
	removeElement(node, 10)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
