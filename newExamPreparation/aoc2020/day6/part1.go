package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	slice := []string{}
	mappa := make(map[string]int)
	num := 0
	for b.Scan() {
		fmt.Println(b.Text())
		if b.Text() != "" {
			slice = append(slice, b.Text())
		} else {
			fmt.Println(len(slice))
			stringa := strings.Join(slice, "")
			// fmt.Println(stringa)
			for _, k := range stringa{
				mappa[string(k)]++
			}

			for _, v := range mappa {
				if v == len(slice) {
					fmt.Println(mappa, v, len(slice), num)
					num++
				}
			}

			// fmt.Println(len(mappa), num)
			slice = []string{}
			mappa = make(map[string]int)
		}
	}
	stringa := strings.Join(slice, "")
	// fmt.Println(stringa)
	for _, k := range stringa{
		mappa[string(k)]++
	}

	for _, v := range mappa {
		if v == len(slice) {
			num++
		}
	}
	fmt.Println(num)
}
