/*

le liste concatenate consistono in una catena di strutture chiamate nodi.
ogni nodo continee un puntatore al prossimo nodo della catena e un valore.
la lista concatenate è una struttura dati che consente di inserire e rimuovere elementi in testa e in coda.
L'ultimo nodo della catena ha un puntatore a nil, che indica la fine della catena.

type Node struct {
	Value int
	Next  *Node
}

*/

package main

//definizione di tipo nodo
type Node struct {
	item int
	next *Node
}

//definizione di tipo lista
type linkedList struct {
	head *Node
}

func newNode(item int) *Node {
	node := new(Node)
	node.item = item
	return node
}

func addNewNode(list linkedList, item int) {
	node := newNode(item)
	node.next = list.head
	list.head = node
}

func main() {

	// creiamo una lista concatenate vuota
	var list linkedList

	// creiamo un nuovo nodo
	var node *Node
	node = new(Node)
	node.item = 10

	// ancora piú veloce
	node = &Node{item: 10, next: nil}

	// inserimento in testa
	var list linkedList
	newNode(30)
	node.next = list.head
	list.head = node

	list.head = node

}
