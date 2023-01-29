package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

type Tree struct {
	Root *Node
}

func newNode(k int) *Node {
	return &Node{Val: k}
}

func insertNode(parent *Node, newNode *Node) {

	if newNode.Val < parent.Val {
		if parent.Right == nil {
			parent.Right = newNode
		} else {
			insertNode(parent.Right, newNode)
		}
	} else {
		if parent.Left == nil {
			parent.Left = newNode
		} else {
			insertNode(parent.Left, newNode)
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

func arr2tree(a []int, i int) (root *Node) {

	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	fmt.Println(root, i)
	root.Left = arr2tree(a[i:], 2*i+1)
	root.Right = arr2tree(a[i:], 2*i+2)
	return root
}

func preorder(node *Node) {
	if node != nil {
		fmt.Println(node.Val)
		preorder(node.Left)
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

	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}

	x := arr2tree(a, 0)

	stampaAlberoASommario(x, 0)

}
