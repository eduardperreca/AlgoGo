package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Head *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func addNewNode(l *Stack, k int) {
	node := newNode(k)
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func push(l *Stack, k int) {
	addNewNode(l, k)
}

func pop(l *Stack) int {
	if l.Head == nil {
		return 0
	}
	value := l.Head.Val
	l.Head = l.Head.Next
	return value
}

func main() {

	stringa := "253-*"

	risultato := 0
	p := &Stack{}
	for _, k := range stringa {
		if unicode.IsDigit(k) {
			numb, _ := strconv.Atoi(string(k))
			push(p, numb)
		} else {
			if k == '-' {
				numb := pop(p)
				numb2 := pop(p)
				risultato = numb2 - numb
				push(p, risultato)
			} else if k == '+' {
				risultato = pop(p) + pop(p)
				push(p, risultato)
			} else if k == '*' {
				risultato = pop(p) * pop(p)
				push(p, risultato)
			} else {
				risultato = pop(p) / pop(p)
				push(p, risultato)
			}
		}
	}
	fmt.Println(risultato)

}
