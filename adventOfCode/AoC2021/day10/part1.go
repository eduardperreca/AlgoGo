package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value string
	Next  *Node
}

type Pile struct {
	Head *Node
}

func newNode(value string) *Node {
	return &Node{Value: value}
}

func addNewNode(pile *Pile, value string) {
	newNode := newNode(value)
	newNode.Next = pile.Head
	pile.Head = newNode
}

func pop(p *Pile) string {
	if p.Head == nil {
		return ""
	}
	value := p.Head.Value
	p.Head = p.Head.Next
	return value
}

func push(p *Pile, value string) {
	addNewNode(p, value)
}

func main() {

	var pile Pile
	// ): 3 points.
	// ]: 57 points.
	// }: 1197 points.
	// >: 25137 points.

	var pile2 Pile

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		check := "{[(<"
		for _, k := range input {
			if strings.Contains(check, string(k)) {
				push(&pile, string(k))
			} else {
				if k == ')' && pile.Head.Value != "(" {
					push(&pile2, string(k))
				} else if k == ']' && pile.Head.Value != "[" {
					push(&pile2, string(k))
				} else if k == '}' && pile.Head.Value != "{" {
					push(&pile2, string(k))
				} else if k == '>' && pile.Head.Value != "<" {
					push(&pile2, string(k))
				} else{
					pop(&pile)
				}
			}
		}
	}
	for pile2.Head != nil {
		fmt.Println(pile2.Head.Value)
		pile2.Head = pile2.Head.Next
	}

}
