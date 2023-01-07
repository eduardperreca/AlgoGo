package main

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func addNodeTree(tree *Tree, value int) {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &Tree{Value: value}
		} else {
			addNodeTree(tree.Left, value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &Tree{Value: value}
		} else {
			addNodeTree(tree.Right, value)
		}
	}
}

func inorder(tree *Tree) {
	if tree == nil {
		return
	}
	inorder(tree.Left)
	println(tree.Value)
	inorder(tree.Right)
}

func preorder(tree *Tree) {
	if tree == nil {
		return
	}
	println(tree.Value)
	preorder(tree.Left)
	preorder(tree.Right)
}

func postorder(tree *Tree) {
	if tree == nil {
		return
	}
	postorder(tree.Left)
	postorder(tree.Right)
	println(tree.Value)
}

func main() {

	b := &Tree{}
	addNodeTree(b, 1)
	addNodeTree(b, 2)
	addNodeTree(b, 3)
	addNodeTree(b, 4)	
	addNodeTree(b, 5)
	addNodeTree(b, 6)
	addNodeTree(b, 7)
	addNodeTree(b, 8)
	addNodeTree(b, 9)
	addNodeTree(b, 10)
	addNodeTree(b, 11)
	
	// inorder
	fmt.Println("Inorder")
	inorder(b)

	fmt.Println("Preorder")
	preorder(b)

	fmt.Println("Postorder")
	postorder(b)

}
