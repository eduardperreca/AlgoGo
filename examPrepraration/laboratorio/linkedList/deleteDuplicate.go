package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
}

func addNode(l *LinkedList, value int) {
	if l.head == nil {
		l.head = &Node{value: value}
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{value: value}
}

func main() {

	l := &LinkedList{}
	addNode(l, 1)
	addNode(l, 2)
	addNode(l, 3)
	addNode(l, 4)
	addNode(l, 4)
	addNode(l, 5)
	for l.head != nil {
		fmt.Println(l.head.value)
		l.head = l.head.next
	}

	fmt.Println("After deleteDuplicate")

	mappa := make(map[int]bool)
	for l.head.next.next != nil {
		fmt.Println(l.head.value)
		if _, ok := mappa[l.head.next.value]; ok {
			l.head = l.head.next.next
		}
		mappa[l.head.next.value] = true
	}
	
	for l.head != nil {
		fmt.Println(l.head.value)
		l.head = l.head.next
	}
	


}
