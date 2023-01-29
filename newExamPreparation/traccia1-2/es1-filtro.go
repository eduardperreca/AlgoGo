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

func f(p *Node, k int) int {
	var a int
	if p == nil {
		return 0
	}
	a = 1 + f(p.Next, k)
	if a == k {
		fmt.Println(p.Val)
	}
	return a
}

func main() {

	l := &linkedList{}
	addNewNode(l, 7)
	addNewNode(l, 1)
	addNewNode(l, 0)
	addNewNode(l, -2)
	addNewNode(l, 3)
	fmt.Println(f(l.Head, 10))
}

// ricevendo il puntatore alla testa di una funzione k
// la funzione f stampa la il valore del nodo di posizione a
// quando é uguale a k, e restituisce a, il valore della lunghezza della lista.

