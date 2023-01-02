package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	i := 0
	numbers := []int{}
	tables := [][]int{}
	for b.Scan() {

		if i == 0 {
			dummy := strings.Split(b.Text(), ",")
			for _, k := range dummy {
				numb, _ := strconv.Atoi(string(k))
				numbers = append(numbers, numb)
			}
		}
		if b.Text() != "" {
			temp := []int{}
			dummy := strings.Split(b.Text(), " ")
			for _, k := range dummy {
				numb, _ := strconv.Atoi(string(k))
				temp = append(temp, numb)
			}
			if len(temp) != 5 && len(temp) != 1 {
				fmt.Println("Errore")
				fmt.Println(temp)
				break
			}
			if len(temp) != 1 {
				tables = append(tables, temp)
			}
		}
		i++
	}

	fmt.Println(numbers, "numberi")
	fmt.Println(tables, "tabelle")

}
