package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func insertNode(parent *Node, newNode *Node) {

	if newNode.Val < parent.Val {
		if parent.Left == nil {
			parent.Left = newNode
		} else {
			insertNode(parent.Left, newNode)
		}
	} else {
		if parent.Right == nil {
			parent.Right = newNode
		} else {
			insertNode(parent.Right, newNode)
		}
	}

}

func addNewNode(t *Tree, k int) {
	node := newNode(k)
	if t.Root == nil {
		t.Root = node
	} else {
		insertNode(t.Root, node)
	}
}

func preorder(node *Node) {
	if node != nil {
		fmt.Println(node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
}

func postorder(node *Node) {
	if node != nil {
		preorder(node.Left)
		preorder(node.Right)
		fmt.Println(node.Val)
	}
}

func inorder(node *Node) {
	if node != nil {
		preorder(node.Left)
		fmt.Println(node.Val)
		preorder(node.Right)
	}
}

func stampaAlberoASommario(node *Node, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.Val)
	if node.Left != nil || node.Right != nil {
		stampaAlberoASommario(node.Right, spaces+1)
		stampaAlberoASommario(node.Left, spaces+1)
	}
}

func main() {

	t := &Tree{}

	addNewNode(t, 3)
	addNewNode(t, 1)
	addNewNode(t, 5)
	addNewNode(t, 10)
	addNewNode(t, 4)

	node := t.Root

	stampaAlberoASommario(node, 0)

	fmt.Println("Preorder")
	preorder(node)

	fmt.Println("Postorder")
	postorder(node)

	fmt.Println("In-Order")
	inorder(node)
}
