package main

type Node struct {
	Val  int
	Next *Node
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

	l := &linkedList{}
	addNewNode(l, 12)
	addNewNode(l, 2)
	addNewNode(l, 4)

}
