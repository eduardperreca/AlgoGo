package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
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

	var l = &linkedList{}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		val := 0
		fmt.Scan(&val)
		addNewNode(l, val)
	}
	fmt.Println("f con 1")
	fmt.Println(f(l.Head, 1))
	fmt.Println("f con 5")
	fmt.Println(f(l.Head, 5))
	fmt.Println("f con 10")
	fmt.Println(f(l.Head, 10))
}
