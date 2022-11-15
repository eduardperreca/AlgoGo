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

	ans := 0
	var pile Pile
	// ): 3 points.
	// ]: 57 points.
	// }: 1197 points.
	// >: 25137 points.

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()

		check := "{[(<"
		fmt.Println(input)
		for _, k := range input {
			if strings.Contains(check, string(k)) {
				push(&pile, string(k))
			}
			if pile.Head.Value == "(" && string(k) != ")" {
				pop(&pile)
				ans += 3
				break
			} else if pile.Head.Value == "[" && string(k) != "]" {
				pop(&pile)
				ans += 57
				break
			} else if pile.Head.Value == "{" && string(k) != "}" {
				pop(&pile)
				ans += 1197
				break
			} else if pile.Head.Value == "<" && string(k) != ">" {
				pop(&pile)
				ans += 25137
				break
			}
		}
		
	}

	fmt.Println(ans)

}
