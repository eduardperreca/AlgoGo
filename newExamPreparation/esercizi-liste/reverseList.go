package main

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

func main() {

	var l = &linkedList{}

	addNewNode(l, 5)
	addNewNode(l, 10)
	addNewNode(l, 12)
	addNewNode(l, 4)
	addNewNode(l, 10)

}
