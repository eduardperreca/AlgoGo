package main

import (
	"bufio"
	"fmt"
	"os"
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

	counter := 0
	stack := []string{}
	
	for b.Scan() {
		for _, k := range b.Text() {
			if values[string(k)] {
				stack = append(stack, string(k))
				fmt.Println(stack)
			} else {
				check := stack[len(stack)-1]
				if string(k) == "}" && check != "{" {
					counter += 1197
					break
				} else if string(k) == "]" && check != "[" {
					counter += 57
					break
				} else if string(k) == ")" && check != "("  {
					counter += 3
					break
				} else if string(k) == ">" && check != "<" {
					counter += 25137
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	fmt.Println(counter)
}
