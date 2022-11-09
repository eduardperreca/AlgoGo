package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value string
	Next  *Node
}

type Pile struct {
	Head *Node
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

func newNode(value string) *Node {
	return &Node{Value: value}
}

func addNewNode(pile *Pile, value string) {
	newNode := newNode(value)
	newNode.Next = pile.Head
	pile.Head = newNode
}

func main() {

	var pile Pile
	var pile2 Pile

	var input = "<a> <b> </b> </c>"

	var listCheck = []string{}

	for i := len(strings.Split(input, " ")); i > 0; i-- {
		listCheck = append(listCheck, strings.Split(input, " ")[i-1])
	}

	for _, line := range listCheck {
		addNewNode(&pile, line)
	}

	for pile.Head != nil {
		if !strings.Contains(pile.Head.Value, "</") {
			push(&pile2, pile.Head.Value)
		} else {
			t := pop(&pile2)
			splitT := strings.Split(t, "<")[1]
			splitP := strings.Split(pile.Head.Value, "</")[1]

			if splitT != splitP {
				fmt.Println(t, pile.Head.Value, "false")
				return
			}
		}
		pile.Head = pile.Head.Next
	}

	if pile2.Head != nil {
		for pile2.Head != nil {
			fmt.Print(pile2.Head.Value, " ")
			pile2.Head = pile2.Head.Next
		}
		fmt.Println("not empty")
	} else {
		fmt.Println("all good")
	}
}
