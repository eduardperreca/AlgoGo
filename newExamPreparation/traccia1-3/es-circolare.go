package main

import "fmt"

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

type linkedList struct {
	Head *Node
	Tail *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *linkedList, k int) {
	node := newNode(k)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Previous = node
		l.Head = node
		l.Head.Previous = l.Tail
		l.Tail.Next = l.Head
	}
}

// func f(p *Node, k int) int {
// 	var a int
// 	if p == nil {
// 		return 0
// 	}
// 	a = 1 + f(p.Next, k)
// 	if a == k {
// 		fmt.Println(p.Val)
// 	}
// 	return a
// }

func f2(p *Node, k int) int {
	var a = 1
	start := p.Previous
	end := p
	for start != end {
		if a == k {
			fmt.Println(start.Val)
		}
		a++
		start = start.Previous
	}

	return a
}

func sizeOf(node *Node) int {
	visited := make(map[*Node]bool)
	for !visited[node] {
		visited[node] = true
		node = node.Next
	}
	return len(visited)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func sposta(p *Node) {
	visited := make(map[*Node]bool)
	for len(visited) != sizeOf(p) {
		if !visited[p] {
			visited[p] = true
			n := p.Val
			n = abs(n)
			for i := 0; i < n; i++ {
				if p.Val < 0 {
					p.Val, p.Previous.Val = p.Previous.Val, p.Val
				} else if p.Val > 0 {
					p.Val, p.Next.Val = p.Next.Val, p.Val
				}
			}
		}
		p = p.Next
	}
}

func main() {
	l := &linkedList{}
	addNewNode(l, 7)
	addNewNode(l, 1)
	addNewNode(l, 0)
	addNewNode(l, -2)
	addNewNode(l, 3)
	// fmt.Println(sizeOf(l.Head), "questo")
	// fmt.Println(f2(l.Head, 1))
	node := l.Head
	for i := 0; i<5; i++{
		fmt.Println(node.Val)
		node = node.Next
	}
	fmt.Println("post changes")
	sposta(l.Head)
	for i := 0; i<5; i++{
		fmt.Println(node.Val)
		node = node.Previous
	}
}
