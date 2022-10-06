package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	t := ""
	var mappa = make(map[string]int)
	b := bufio.NewScanner(os.Stdin)
	for b.Scan() {
		if b.Text() == "exit" {
			break
		}
		for _, k := range b.Text() {
			if string(k) >= "a" && string(k) <= "z" {
				t = string(k - 32)
			} else if string(k) >= "A" && string(k) <= "Z" {
				t = string(k)
			} else {
				continue
			}
			mappa[t]++
		}
		fmt.Println(mappa)
		array := []string{}
		for k := range mappa {
			array = append(array, k)
		}
		sort.Strings(array)
		for _, k := range array {
			dummy := ""
			for i := 0; i < mappa[k]; i++ {
				dummy += "*"
			}
			fmt.Println(k, dummy)
		}
	
	}
}
