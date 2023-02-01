package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func in(mappa map[string][]string, name string) bool {
	for k := range mappa {
		if k == name {
			return true
		}
	}
	return false
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	mappa := make(map[string][]string)
	for b.Scan() {
		parents := strings.Split(b.Text(), "contain")
		parent := strings.ReplaceAll(parents[0], "bags", "")
		parent = strings.TrimRight(parent, " ")
		parent = strings.TrimLeft(parent, " ")
		// fmt.Println(parent)

		bags := strings.Split(parents[1], ", ")
		for _, bag := range bags {
			if bag != " no other bags." || bag != "other bags." {
				bag = strings.TrimLeft(bag, " ")
				bag = strings.TrimRight(bag, " ")
				bagSlice := strings.Split(bag, " ")
				bagName := strings.Join(bagSlice[1:3], " ")
				mappa[bagName] = append(mappa[bagName], parent)
				fmt.Println(bagName)
			}
		}
	}
	answer := []string{}
	dfs := []string{"shiny gold"}
	visited := make(map[string]bool)
	for len(dfs) > 0 {
		el := dfs[len(dfs)-1]
		dfs = dfs[:len(dfs)-1]
		fmt.Println(el, dfs)
		if !visited[el] {
			visited[el] = true
			if in(mappa, el) {
				parents := mappa[el]
				for _, p := range parents {
					dfs = append(dfs, p)
				}
			} else {
				answer = append(answer, el)
			}
		}
	}
	fmt.Println(len(visited) - 1)
}
