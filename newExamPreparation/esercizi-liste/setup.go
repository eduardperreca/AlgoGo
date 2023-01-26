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
	fmt.Println(a, "a")
	if a == k {
		fmt.Println(p.Val, "valore")
	}

	return a
}
// ricevendo il puntatore alla testa di una lista e un intero k la funzione f stampa il valore del nodo p quando il valore del target é
// uguale al valore di a, incrementalmente indica la lunghezza della lista concatenata, restituendo in fine la lunghezza della lista.
// Casi limite da analizzare

func main() {

	l := &linkedList{}
	// fmt.Println("Inserisci quanti nodi aggiungere")
	// fmt.Scan(&n)
	// x := 0
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&x)
	// 	addNewNode(l, x)
	// }

	addNewNode(l, 7)
	addNewNode(l, 1)
	addNewNode(l, 0)
	addNewNode(l, -2)
	addNewNode(l, 3)

	z := f(l.Head, 10)
	fmt.Println(z, "z")
}
