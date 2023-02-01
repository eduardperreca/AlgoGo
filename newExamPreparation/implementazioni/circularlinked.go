package main

type Node struct {
	Previous *Node
	Next     *Node
	Val      int
}

type circularList struct {
	Head *Node
	Tail *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *circularList, k int) {
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

	l := &circularList{}
	addNewNode(l, 10)
	addNewNode(l, 12)
	addNewNode(l, 20)
	addNewNode(l, 25)

}
