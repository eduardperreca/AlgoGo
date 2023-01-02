package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var mappa = make(map[string]int)

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	for b.Scan() {
		splitted := strings.Split(b.Text(), " -> ")
		fmt.Println(splitted[0], splitted[1])

		splittedX := strings.Split(splitted[0], ",")
		splittedX2 := strings.Split(splitted[1], ",")
		numb, _ := strconv.Atoi(splittedX[0])
		numb2, _ := strconv.Atoi(splittedX[1])
		numb3, _ := strconv.Atoi(splittedX2[0])
		numb4, _ := strconv.Atoi(splittedX2[1])

	
		
		// fmt.Println(splittedX[0], splittedX[1])
		// fmt.Println(splittedX2[0], splittedX2[1])

	}
}
