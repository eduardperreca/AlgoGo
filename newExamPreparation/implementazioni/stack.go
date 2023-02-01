package main

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

func addNewNode(s *Stack, k int) {
	node := newNode(k)
	if s.Head == nil {
		s.Head = node
	} else {
		node.Next = s.Head
		s.Head = node
	}
}

func push(s *Stack, k int) {
	addNewNode(s, k)
}

func pop(s *Stack) int {
	if s.Head == nil {
		return 0
	}
	val := s.Head.Val
	s.Head = s.Head.Next
	return val
}

func main() {

}
