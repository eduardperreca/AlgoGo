package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}

func addNewNode(list *List, value int) {
	node := newNode(value)
	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
}

// func sortingList(list *List) {
// 	var node *Node
// 	var node2 *Node
// 	for node = list.Head; node != nil; node = node.Next {
// 		for node2 = node.Next; node2 != nil; node2 = node2.Next {
// 			if node.Value > node2.Value {
// 				node.Value, node2.Value = node2.Value, node.Value
// 			}
// 		}
// 	}
// }

func printList(l *List) {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

func main() {

	var l List
	addNewNode(&l, 1)
	addNewNode(&l, 4)
	addNewNode(&l, 12)
	addNewNode(&l, 8)
	addNewNode(&l, 6)
	addNewNode(&l, 32)
	addNewNode(&l, 9)
	printList(&l)
	fmt.Println("Sorting...")
	printList(&l)
	fmt.Println("Done!")

}
