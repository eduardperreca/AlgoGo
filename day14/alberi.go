package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

type Tree struct {
	Head *Node
}

func arr2tree(a []int, i int) (root *Node) {
	if i >= len(a) {
		return nil
	}
	root = &Node{Value: a[i]}
	root.Left = arr2tree(a, 2*i+1)
	root.Right = arr2tree(a, 2*i+2)
	return root
}

func printTree(t *Tree) {
	if t.Head == nil {
		return
	}
	fmt.Println(t.Head.Value)
	printTree(&Tree{t.Head.Left})
	printTree(&Tree{t.Head.Right})
}

func printPostOrder(t *Tree) {
	if t.Head == nil {
		return
	}
	printPostOrder(&Tree{t.Head.Left})
	printPostOrder(&Tree{t.Head.Right})
	fmt.Println(t.Head.Value)
}

func printInOrder(t *Tree) {
	if t.Head == nil {
		return
	}
	printInOrder(&Tree{t.Head.Left})
	fmt.Println(t.Head.Value)
	printInOrder(&Tree{t.Head.Right})
}

func stampaAlberoASommario(a *Node, level int) {
	if a == nil {
		return
	}
	stampaAlberoASommario(a.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(a.Value)
	stampaAlberoASommario(a.Left, level+1)
}

func stampaAlbero(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, "")
	if node.Left == nil && node.Right == nil {
		return
	}
	fmt.Print(" [")
	if node.Left != nil {
		stampaAlbero(node.Left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.Right != nil {
		stampaAlbero(node.Right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

func main() {

	i := 0
	const N = 10

	var slice = make([]int, 0, N)

	for i < N {
		slice = append(slice, rand.Intn(99))
		i++
	}

	tree := arr2tree(slice, 0)

	// fmt.Println("print normale")
	// printTree(&Tree{tree})

	// fmt.Println("inposta")
	// printInOrder(&Tree{tree})

	// fmt.Println("postposta")
	// printPostOrder(&Tree{tree})

	fmt.Println("albero a sommario")
	stampaAlberoASommario(tree, 0)

	fmt.Println("albero")
	stampaAlbero(tree)

}
