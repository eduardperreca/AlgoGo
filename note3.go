// implemeantazione alberi binari pazzeschi

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

// ---------------------------------------------- //

// implementazione stack senza slices

package main

import (
	"fmt"
	"strconv"
)

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

func pop(p *Pile) int {
	if p.Head == nil {
		return 0
	}
	value := p.Head.Value
	p.Head = p.Head.Next
	return value
}

func push(p *Pile, value int) {
	addNewNode(p, value)
}

func isEmpty(p *Pile) bool {
	return p.Head == nil
}

func size(list *Pile) int {
	var size int
	for node := list.Head; node != nil; node = node.Next {
		size++
	}
	return size
}

func main() {

	var p Pile
	numbers := []string{"5", "3", "-", "2", "*"}

	for _, k := range numbers {
		if k == "+" {
			push(&p, pop(&p)+pop(&p))
		} else if k == "-" {
			a := pop(&p)
			b := pop(&p)
			push(&p, b-a)
		} else if k == "*" {
			push(&p, pop(&p)*pop(&p))
		} else if k == "/" {
			a := pop(&p)
			b := pop(&p)
			push(&p, b/a)
		} else {
			numb, _ := strconv.Atoi(k)
			push(&p, numb)
		}
	}

	curr := p.Head

	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}

}



// ---------------------------------------------- //


// sorting linked list
package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}

func addNewNode(list *List, value int) {
	node := newNode(value)
	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
}

// func sortingList(list *List) {
// 	var node *Node
// 	var node2 *Node
// 	for node = list.Head; node != nil; node = node.Next {
// 		for node2 = node.Next; node2 != nil; node2 = node2.Next {
// 			if node.Value > node2.Value {
// 				node.Value, node2.Value = node2.Value, node.Value
// 			}
// 		}
// 	}
// }

func printList(l *List) {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

func main() {

	var l List
	addNewNode(&l, 1)
	addNewNode(&l, 4)
	addNewNode(&l, 12)
	addNewNode(&l, 8)
	addNewNode(&l, 6)
	addNewNode(&l, 32)
	addNewNode(&l, 9)
	printList(&l)
	fmt.Println("Sorting...")
	printList(&l)
	fmt.Println("Done!")

}


// ---------------------------------------------- //

// aoc lanternfish 
package main

import (
	"fmt"
)

func main() {

	toParse := []int{
		4, 1, 4, 1, 3, 3, 1, 4, 3, 3, 2, 1, 1, 3, 5, 1, 3, 5, 2, 5, 1, 5, 5, 1, 3, 2, 5, 3,
		1, 3, 4, 2, 3, 2, 3, 3, 2, 1, 5, 4, 1, 1, 1, 2, 1, 4, 4, 4, 2, 1, 2, 1, 5, 1, 5, 1,
		2, 1, 4, 4, 5, 3, 3, 4, 1, 4, 4, 2, 1, 4, 4, 3, 5, 2, 5, 4, 1, 5, 1, 1, 1, 4, 5, 3,
		4, 3, 4, 2, 2, 2, 2, 4, 5, 3, 5, 2, 4, 2, 3, 4, 1, 4, 4, 1, 4, 5, 3, 4, 2, 2, 2, 4,
		3, 3, 3, 3, 4, 2, 1, 2, 5, 5, 3, 2, 3, 5, 5, 5, 4, 4, 5, 5, 4, 3, 4, 1, 5, 1, 3, 4,
		4, 1, 3, 1, 3, 1, 1, 2, 4, 5, 3, 1, 2, 4, 3, 3, 5, 4, 4, 5, 4, 1, 3, 1, 1, 4, 4, 4,
		4, 3, 4, 3, 1, 4, 5, 1, 2, 4, 3, 5, 1, 1, 2, 1, 1, 5, 4, 2, 1, 5, 4, 5, 2, 4, 4, 1,
		5, 2, 2, 5, 3, 3, 2, 3, 1, 5, 5, 5, 4, 3, 1, 1, 5, 1, 4, 5, 2, 1, 3, 1, 2, 4, 4, 1,
		1, 2, 5, 3, 1, 5, 2, 4, 5, 1, 2, 3, 1, 2, 2, 1, 2, 2, 1, 4, 1, 3, 4, 2, 1, 1, 5, 4,
		1, 5, 4, 4, 3, 1, 3, 3, 1, 1, 3, 3, 4, 2, 3, 4, 2, 3, 1, 4, 1, 5, 3, 1, 1, 5, 3, 2,
		3, 5, 1, 3, 1, 1, 3, 5, 1, 5, 1, 1, 3, 1, 1, 1, 1, 3, 3, 1
	}

	var mappa = make(map[int]int)
	for _, v := range toParse {
		mappa[v]++
	}
	for i := 0; i < 256; i++ {
		mappa2 := make(map[int]int)
		for k, v := range mappa {
			if k == 0 {
				mappa2[6] += v
				mappa2[8] += v
			} else {
				mappa2[k-1] += v
			}
		}
		mappa = mappa2

		fmt.Println(mappa2, " mappa2 all' iterazione ", i)
	}
	val := 0
	for _, k := range mappa {
		val += k
	}
	fmt.Println(val)
}


// ---------------------------------------------- //
