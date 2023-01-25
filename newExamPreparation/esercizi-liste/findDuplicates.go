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

func findDuplicates(node *Node, k int) int {
	var mappa = make(map[int]bool)
	for node != nil {
		if !mappa[node.Val] {
			mappa[node.Val] = true
		} else {
			return node.Val
		}
		node = node.Next
	}
	return -1
}

func main() {

	l := &linkedList{}
	addNewNode(l, 5)
	addNewNode(l, 10)
	addNewNode(l, 12)
	addNewNode(l, 4)
	addNewNode(l, 10)

	node := l.Head
	t := findDuplicates(node, 10)

	if t > 0 {
		fmt.Println("found duplicate", t)
	} else {
		fmt.Println("not found")
	}

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}
