package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func arr2tree(a []int, i int) (root *Node) {
	if i >= len(a) {
		return nil
	}

	root = &Node{Val: a[i]}
	root.Left = arr2tree(a, 2*i+1)
	root.Right = arr2tree(a, 2*i+2)
	return
}

// func stampaAlberoASommario(node *Node, spaces int) {
// 	for i := 0; i < spaces; i++ {
// 		fmt.Print(" ")
// 	}
// 	fmt.Print("*")
// 	fmt.Println(node.Val)
// 	if node.Left != nil || node.Right != nil {
// 		stampaAlberoASommario(node.Right, spaces+1)
// 		stampaAlberoASommario(node.Left, spaces+1)
// 	}
// }

func preorder(root *Node){
	if root == nil{
		return
	}
	fmt.Println(root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	root := arr2tree(array, 0)
	preorder(root)
}
