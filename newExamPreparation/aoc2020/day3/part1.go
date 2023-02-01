package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkTrees(matrix [][]string, stepi int, stepj int) int {
	idk := 0
	i, j := 0, 0
	for i < len(matrix) {
		if matrix[i][j] == "#" {
			idk++
		}
		i += stepi
		j += stepj
		if j >= len(matrix[0]) {
			j -= len(matrix[0])
		}
	}
	return idk
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	matrix := [][]string{}
	for b.Scan() {
		sliceToMatrix := []string{}
		for _, k := range b.Text() {
			sliceToMatrix = append(sliceToMatrix, string(k))
		}
		matrix = append(matrix, sliceToMatrix)
	}
	fmt.Println(len(matrix))
	fmt.Println(len(matrix[0]))

	var n int
	fmt.Scan(&n)
	mult := 1
	for i := 0; i < n; i++ {
		var stepi, stepj int
		fmt.Scan(&stepi, &stepj)
		mult *= checkTrees(matrix, stepi, stepj)
	}
	fmt.Println(mult)
}
