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

func reverseList(head *Node) *Node {
	var prev *Node
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
		fmt.Println(prev, head, head.Next)
		fmt.Println(head, head.Next, prev)
	}
	return prev
}

func main() {

	var l = &linkedList{}

	addNewNode(l, 5)
	addNewNode(l, 10)
	addNewNode(l, 12)
	addNewNode(l, 4)
	addNewNode(l, 10)

	node := l.Head
	reverseList(node)

}
