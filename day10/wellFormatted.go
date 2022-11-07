package main

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

func main(){

	listCheck := []string{"<a>", "<b>", "</b>", "</a>"}


}