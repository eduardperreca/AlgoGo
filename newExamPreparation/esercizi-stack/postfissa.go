package main

import (
	"fmt"
	"unicode"
)

type Stack struct {
	Head *Node
}

type Node struct {
	Next *Node
	Val  string
}

func addnewNode(s *Stack, k string) {
	node := newNode(k)
	if s.Head == nil {
		s.Head = node
	} else {
		node.Next = s.Head
		s.Head = node
	}
}

func push(s *Stack, k string) {
	addnewNode(s, k)
}

func pop(s *Stack) string {
	if s.Head == nil {
		return ""
	}
	value := s.Head.Val
	s.Head = s.Head.Next
	return value
}

func newNode(k string) *Node {
	return &Node{Val: k}
}

func main() {

	stringa := "((5-3)*2)"
	values := map[string]bool{
		"*": true,
		"-": true,
		"+": true,
		"/": true,
	}
	l := &Stack{}
	for _, k := range stringa {
		if values[string(k)]{
			push(l, string(k))
		}
		if unicode.IsDigit(k) {
			fmt.Print(string(k))
		}
		if k == ')'{
			fmt.Print(string(pop(l)))
		}
	}

}
