package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	values := map[string]bool{
		"{": true,
		"[": true,
		"(": true,
		"<": true,
	}

	valuestoAdd := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	stackTotal := [][]string{}
	stack := []string{}
	for b.Scan() {
		buono := true
		stack = []string{}
		for _, k := range b.Text() {
			if values[string(k)] {
				stack = append(stack, string(k))
				// fmt.Println(stack)
			} else {
				check := stack[len(stack)-1]
				if string(k) == "}" && check != "{" {
					buono = false
					break
				} else if string(k) == "]" && check != "[" {
					buono = false
					break
				} else if string(k) == ")" && check != "(" {
					buono = false
					break
				} else if string(k) == ">" && check != "<" {
					buono = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
		if buono {
			stackTotal = append(stackTotal, stack)
		}
	}

	// fmt.Println(stackTotal)

	totVet := []int{}
	tot := 0
	for _, k := range stackTotal {
		tot = 0
		for i := len(k) - 1; i >= 0; i-- {
			tot = tot*5 + valuestoAdd[k[i]]
		}
		totVet = append(totVet, tot)
	}
	// fmt.Println(totVet)
	fmt.Println(len(stackTotal))
	fmt.Println(len(totVet))
	sort.Ints(totVet)
	fmt.Println(totVet)
	fmt.Println(totVet[len(totVet)/2])
}
