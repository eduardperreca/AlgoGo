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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *Tree) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func printTree(tree *Tree) {
	if tree == nil {
		return
	}
	printTree(tree.Left)
	fmt.Println(tree.Value)
	printTree(tree.Right)
}

func main() {

	tree := &Tree{Value: 10}
	addNodeTree(tree, 5)
	addNodeTree(tree, 15)
	addNodeTree(tree, 3)
	addNodeTree(tree, 7)
	addNodeTree(tree, 12)
	addNodeTree(tree, 17)
	addNodeTree(tree, 1)
	addNodeTree(tree, 4)
	addNodeTree(tree, 6)
	addNodeTree(tree, 8)
	addNodeTree(tree, 11)
	addNodeTree(tree, 13)
	addNodeTree(tree, 16)
	addNodeTree(tree, 18)

	printTree(tree)

}
