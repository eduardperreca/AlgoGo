package main

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

func main() {

	

}
