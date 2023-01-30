package main

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

func insertNode(parent *Node, node *Node) {

	if node.Val < parent.Val {
		if parent.Left == nil {
			parent.Left = node
		} else {
			insertNode(parent.Left, node)
		}
	} else {
		if parent.Right == nil {
			parent.Right = node
		} else {
			insertNode(parent.Right, node)
		}
	}

}

func addNewNode(root *Tree, k int) {
	node := newNode(k)
	if root.Root == nil {
		root.Root = node
	} else {
		insertNode(root.Root, node)
	}
}

func main() {

}
