package main

import "fmt"

type Node struct {
	Row    int
	Column int
	Val    int
}

func min(a []Node, target int) Node {
	min := a[0].Val
	for _, v := range a {
		if v.Val == target {
			return v
		}
	}

	for _, v := range a {
		if v.Val < min {
			min = v.Val
		}
	}
	for _, v := range a {
		if v.Val == min {
			return v
		}
	}
	return Node{}
}

func main() {
	matrix := [][]int{
		{7, 1, 8, 4, 4},
		{9, 1, 8, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 1},
	}

	// 1. Create a stack
	stack := []Node{}

	// 2. Add the first node to the stack
	stack = append(stack, Node{0, 0, 0})
	dummy := []Node{}
	// 3. Create visited map
	visited := make(map[Node]bool)
	count := 0
	valSum := 0
	// 4. While the stack is not empty
	for len(stack) > 0 {

		// 5. Pop the first node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node)
		fmt.Println("value:", matrix[node.Row][node.Column])

		// 6. Check if it is the goal
		if node.Column == len(matrix)-1 && node.Row == len(matrix[0])-1 {
			fmt.Println("END!, we spent", count, "steps")
			fmt.Println("END!, we spent", valSum, "value")
			return
		}
		dummy = []Node{}
		// 7. Add the neighbors to the stack
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column, matrix[node.Row-1][node.Column]}]; !ok {
				visited[Node{node.Row - 1, node.Column, matrix[node.Row-1][node.Column]}] = true
				dummy = append(dummy, Node{node.Row - 1, node.Column, matrix[node.Row-1][node.Column]})
			}
		}

		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1, matrix[node.Row][node.Column-1]}]; !ok {
				visited[Node{node.Row, node.Column - 1, matrix[node.Row][node.Column-1]}] = true
				dummy = append(dummy, Node{node.Row, node.Column - 1, matrix[node.Row][node.Column-1]})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1, matrix[node.Row][node.Column+1]}]; !ok {
				visited[Node{node.Row, node.Column + 1, matrix[node.Row][node.Column+1]}] = true
				dummy = append(dummy, Node{node.Row, node.Column + 1, matrix[node.Row][node.Column+1]})
			}
		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column, matrix[node.Row+1][node.Column]}]; !ok {
				visited[Node{node.Row + 1, node.Column, matrix[node.Row+1][node.Column]}] = true
				dummy = append(dummy, Node{node.Row + 1, node.Column, matrix[node.Row+1][node.Column]})
			}
		}

		stack = append(stack, min(dummy, matrix[len(matrix)-1][len(matrix[0])-1]))
		fmt.Println("dummy:", dummy)
		count++
		valSum += stack[0].Val

	}
}
