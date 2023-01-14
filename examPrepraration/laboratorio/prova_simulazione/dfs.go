package main

import "fmt"

type Node struct {
	Row    int
	Column int
}

func main() {

	matrix := [][]int{
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 2},
	}

	// 1. Create a stack
	stack := []Node{}

	// 2. Add the first node to the stack
	stack = append(stack, Node{0, 0})

	// 3. Create visited map
	visited := make(map[Node]bool)

	// 4. While the stack is not empty
	for len(stack) > 0 {

		// 5. Pop the first node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node)
		fmt.Println("Stack:", stack)
		fmt.Println("Visited:", visited)

		// 6. Check if it is the goal
		if matrix[node.Row][node.Column] == 2 {
			fmt.Println("END!, we spent", len(visited), "steps")
			return
		}

		// 7. Add the neighbors to the stack
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column}]; !ok && matrix[node.Row-1][node.Column] == 0 || matrix[node.Row-1][node.Column] == 2 {
				visited[Node{node.Row - 1, node.Column}] = true
				stack = append(stack, Node{node.Row - 1, node.Column})
			}
		}

		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1}]; !ok && matrix[node.Row][node.Column-1] == 0 || matrix[node.Row][node.Column-1] == 2 {
				visited[Node{node.Row, node.Column - 1}] = true
				stack = append(stack, Node{node.Row, node.Column - 1})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1}]; !ok && matrix[node.Row][node.Column+1] == 0 || matrix[node.Row][node.Column+1] == 2 {
				visited[Node{node.Row, node.Column + 1}] = true
				stack = append(stack, Node{node.Row, node.Column + 1})
			}
		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column}]; !ok && matrix[node.Row+1][node.Column] == 0 || matrix[node.Row+1][node.Column] == 2 {
				visited[Node{node.Row + 1, node.Column}] = true
				stack = append(stack, Node{node.Row + 1, node.Column})
			}
		}

	}

}
