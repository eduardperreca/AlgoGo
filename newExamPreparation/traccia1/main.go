package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *LinkedList, val int) {
	node := newNode(val)
	node.Next = l.Head
	l.Head = node
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

	n := 0
	fmt.Println("Inserisci quanti nodi vuoi aggiungere: ")
	fmt.Scan(&n)
	l := &LinkedList{}
	k := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		addNewNode(l, k)
	}

	p := l.Head
	z := f(p, 10)
	
	fmt.Println(z)

}
