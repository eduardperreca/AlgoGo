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

func insert(parent *Node, node *Node) {
	if node.Val > parent.Val {
		if parent.Right == nil {
			parent.Right = node
		} else {
			insert(parent.Right, node)
		}
	} else {
		if parent.Left == nil {
			parent.Left = node
		} else {
			insert(parent.Left, node)
		}
	}
}

func addNewNode(r *Tree, k int) {
	node := newNode(k)
	if r.Root == nil {
		r.Root = node
	} else {
		insert(r.Root, node)
	}
}

func arr2tree(a []int, i int) (root *Node) {
	if i >= len(a) {
		return nil
	}
	root = &Node{Val: a[i]}
	root.Right = arr2tree(a, 2*i+1)
	root.Left = arr2tree(a, 2*i+2)
	return
}

func main() {

}
