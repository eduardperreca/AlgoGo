package main

import "fmt"

func main() {

	toParse := []int{3, 4, 3, 1, 2}
	var mappa = make(map[int]int)
	i := 0
	check := 0
	for {
		if i == len(toParse) {
			i = 0
			check++
		}

		if check == 191 {
			fmt.Println(len(toParse))
			break
		}

		toParse[i] = toParse[i] - 1
		if toParse[i] == -1 {
			toParse[i] = 6
			toParse = append(toParse, 9)
		}
		i++
	}

}
