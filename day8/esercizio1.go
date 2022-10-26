package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func addNewNode(list *List, value int) {
	// in questo modo aggiungo un nuovo nodo alla lista
	newNode := newNode(value)
	newNode.Next = list.Head
	list.Head = newNode
}

func newNode(value int) *Node {
	// in questo modo ritorno un puntatore a un nodo
	return &Node{Value: value}
}

func searchNode(list *List, value int) int {
	for node := list.Head; node != nil; node = node.Next {
		if node.Value == value {
			return 1
		}
	}
	return 0
}

func deleteAll(l *List) {
	l.Head = nil
}

func removeFirstNode(list *List) {
	list.Head = list.Head.Next
}

func removeLastNode(list *List) {
	for node := list.Head; node != nil; node = node.Next {
		if node.Next.Next == nil {
			node.Next = nil
		}
	}
}

func printList(list *List) {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

func reversePrintList(list *List) {
	if list.Head == nil {
		return
	}
	reversePrintList(&List{list.Head.Next})
	fmt.Println(list.Head.Value)
}

func main() {

	var l List

	printList(&l)
	fmt.Print("lista prima\n")

	addNewNode(&l, 1)
	addNewNode(&l, 2)
	addNewNode(&l, 3)
	addNewNode(&l, 4)
	addNewNode(&l, 5)

	reversePrintList(&l)
	fmt.Print("lista reverted\n")

	printList(&l)
	fmt.Print("lista dopo")
	fmt.Println()
	removeFirstNode(&l)
	printList(&l)
	fmt.Print("lista pulita del primo nodo")

	fmt.Println()
	removeLastNode(&l)
	printList(&l)
	fmt.Print("lista pulita dell'ultimo nodo")

	fmt.Println()
	deleteAll(&l)
	printList(&l)
	fmt.Print("lista pulita")

}
