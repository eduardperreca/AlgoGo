package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var pile []string
	var abs []int
	for scanner.Scan() {
		input := scanner.Text()
		check := "{[(<"
		pile = []string{}
		for _, k := range input {
			if strings.Contains(check, string(k)) {
				pile = append(pile, string(k))
			} else {
				if pile[len(pile)-1] == "(" && string(k) != ")" {
					continue
				} else if pile[len(pile)-1] == "[" && string(k) != "]" {
					continue
				} else if pile[len(pile)-1] == "{" && string(k) != "}" {
					continue
				} else if pile[len(pile)-1] == "<" && string(k) != ">" {
					continue
				} else {
					pile = pile[:len(pile)-1]
				}
			}
		}
		// fmt.Println(pile)
		dummy := ""
		tot := 0
		for i := len(pile) - 1; i >= 0; i-- {
			if pile[i] == "(" {
				dummy += ")"
				tot = 5*tot + 1
			} else if pile[i] == "[" {
				dummy += "]"
				tot = 5*tot + 2
			} else if pile[i] == "{" {
				dummy += "}"
				tot = 5*tot + 3
			} else if pile[i] == "<" {
				dummy += ">"
				tot = 5*tot + 4
			}
			// fmt.Println(tot)
		}
		abs = append(abs, tot)
		fmt.Println(dummy)

	}
	sort.Ints(abs)
	fmt.Println(len(abs))
	fmt.Println(abs[(len(abs)/2)])
	// for pile.Head != nil {
	// 	fmt.Print(pile.Head.Value)
	// 	if pile.Head.Value == "(" {
	// 		dummy += ")"
	// 	} else if pile.Head.Value == "[" {
	// 		dummy += "]"
	// 	} else if pile.Head.Value == "{" {
	// 		dummy += "}"
	// 	} else if pile.Head.Value == "<" {
	// 		dummy += ">"
	// 	}
	// 	pile.Head = pile.Head.Next
	// }
	// fmt.Println()
	// fmt.Print(dummy)
}
