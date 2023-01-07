package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:n-1]))
}

func main() {
	numbers := []int{1, 2, 5, 7, -2, 10, 9, 21, 3, 8}
	println(largest(numbers))
}
