package main

import "fmt"

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

type circList struct {
	Head *Node
	Tail *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *circList, k int) {
	node := newNode(k)
	node.Next = l.Head
	l.Head = node
}

func addNewBack(l *circList, k int){
	node := newNode(k)
	node.Previous = l.Tail
	l.Tail = node
}

func main() {

	l := &circList{}
	a := 0
	k := 0
	fmt.Println("Inserisci quanti nodi vuoi aggiungere: ")
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&k)
		addNewNode(l, k)
	}

	// var visited = make(map[int]bool)
	p := l.Head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

}
