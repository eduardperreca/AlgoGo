package main

import "fmt"

func main() {

	var stringa string
	fmt.Scan(&stringa)
	var mappa = make(map[string]int)
	for _, k := range stringa {
		mappa[string(k)]++
	}

	if len(mappa)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}

}
