package main

import "fmt"

type Node struct {
	Row    int
	Column int
}

func main() {

	matrix := [][]int{
		{7, 1, 8, 4, 4},
		{9, 6, 22, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 12},
	}

	// 1. Create a queue
	queue := []Node{}

	// 2. Add the first node to the queue
	queue = append(queue, Node{0, 0})
	visited := make(map[Node]bool)
	// 3. While the queue is not empty
	for len(queue) > 0 {

		// 4. Dequeue the first node
		node := queue[0]
		queue = queue[1:]

		fmt.Println("visiting", node)
		fmt.Println("with value: ", matrix[node.Row][node.Column])
		fmt.Println("visited", visited)

		// 5. Check if it is the goal
		if matrix[node.Row][node.Column] == 6 {
			println("Found!")
			return
		}

		// 6. Add the neighbors to the queue
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column}]; !ok {
				visited[Node{node.Row - 1, node.Column}] = true
				queue = append(queue, Node{node.Row - 1, node.Column})
			}

		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column}]; !ok {
				visited[Node{node.Row + 1, node.Column}] = true
				queue = append(queue, Node{node.Row + 1, node.Column})
			}
		}
		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1}]; !ok {
				visited[Node{node.Row, node.Column - 1}] = true
				queue = append(queue, Node{node.Row, node.Column - 1})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1}]; !ok {
				visited[Node{node.Row, node.Column + 1}] = true
				queue = append(queue, Node{node.Row, node.Column + 1})
			}
		}
		fmt.Println("queue", queue)
	}

}
