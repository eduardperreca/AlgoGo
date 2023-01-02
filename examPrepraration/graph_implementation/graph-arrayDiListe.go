/* implementazione di grafi come vettore di liste di adiacenza - implementazione "alla lettera"

// L'insieme dei vertici (individuati dagli interi 0,1,... n-1) è definito implicitamente da una variabile intera n
*/

package main

import (
	"fmt"
)

// grafo rappresentato tramite slice di liste concatenate

type graph struct {
	n         int // numero di vertici
	adiacenti []linkedList
}

func nuovoGrafo(n int) *graph {
	var g graph
	g.n = n
	g.adiacenti = make([]linkedList, n)
	stampaGrafo(&g)
	fmt.Println("finito graphNew")
	return &g
}

/* legge da standard input un grafo */
// se ci sono errori restituisce false
// assumo che le mappe per vertici e archi siano già state create
//FORMATO
// n
// v1 w1
// v2 w2
// v3 w3
// ...
// vm wm
func leggiGrafo() (g *graph) {
	//fmt.Println("Inserisci il numero di vertici")
	var n int
	fmt.Scan(&n)
	g = nuovoGrafo(n)

	fmt.Println("Inserisci gli archi, uno per riga, nel formato v1 v2")
	var v1, v2 int
	_, err := fmt.Scan(&v1, &v2)

	for err == nil {
		//fmt.Println("letti vertici:", v1, v2)
		addNewNode(&g.adiacenti[v1], v2)
		//addNewNode(&g.adiacenti[v2], v1)
		//stampaGrafo(g)
		_, err = fmt.Scan(&v1, &v2)
	}
	return g
}

func stampaGrafo(g *graph) {
	fmt.Printf("Il grafo ha %d nodi\n", g.n)
	for i, lista := range g.adiacenti {
		fmt.Print(i, ": ")
		printList(&lista)
	}
}

// liste concatenate

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func newNode(n int) *listNode {
	return &listNode{n, nil}
}
func addNewNode(l *linkedList, n int) {
	node := newNode(n)
	node.next = l.head
	l.head = node
}

func printList(l *linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item)
		if p.next != nil {
			fmt.Print(" ")
		}
		p = p.next
	}
	fmt.Println()
}

// restituisce il puntatore
func searchList(l linkedList, n int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == n {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func removeItem(l *linkedList, n int) bool {
	var curr, prev *listNode = l.head, nil
	tf := false
	for curr != nil {
		if curr.item == n {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			tf = true
		}
		prev = curr
		curr = curr.next
	}
	return tf
}

func main() {
	g := leggiGrafo()
	stampaGrafo(g)

}
