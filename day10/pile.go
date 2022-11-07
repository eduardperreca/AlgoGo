package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

type Pile struct {
	Head *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}

func addNewNode(pile *Pile, value int) {
	newNode := newNode(value)
	newNode.Next = pile.Head
	pile.Head = newNode
}

func pop(p *Pile) int {
	if p.Head == nil {
		return 0
	}
	value := p.Head.Value
	p.Head = p.Head.Next
	return value
}

func push(p *Pile, value int) {
	addNewNode(p, value)
}

func isEmpty(p *Pile) bool {
	return p.Head == nil
}

func size(list *Pile) int {
	var size int
	for node := list.Head; node != nil; node = node.Next {
		size++
	}
	return size
}

func main() {

	var p Pile
	numbers := []string{"5", "3", "-", "2", "*"}

	for _, k := range numbers {
		if k == "+" {
			push(&p, pop(&p)+pop(&p))
		} else if k == "-" {
			a := pop(&p)
			b := pop(&p)
			push(&p, b-a)
		} else if k == "*" {
			push(&p, pop(&p)*pop(&p))
		} else if k == "/" {
			a := pop(&p)
			b := pop(&p)
			push(&p, b/a)
		} else {
			numb, _ := strconv.Atoi(k)
			push(&p, numb)
		}
	}

	curr := p.Head

	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}

}
