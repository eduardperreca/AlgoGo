package main

import (
	"fmt"
	"os"
	"strconv"
)

// func recursiveDaysLength(slice []int, check int) int {
// 	if check == 0 {
// 		return len(slice)
// 	}
// 	test := 0
// 	for i := 0; i < len(slice); i++ {
// 		slice[i] = slice[i] - 1
// 		if slice[i] == -1 {
// 			slice[i] = 6
// 			test++
// 		}
// 	}
// 	if test > 0 {
// 		for k := 0; k < test; k++ {
// 			slice = append(slice, 8)
// 		}
// 	}
// 	// fmt.Println(slice)

// 	return recursiveDaysLength(slice[:len(slice)/2 +1], check-1)
// }

func main() {

	toParse := []int{3, 4, 3, 1, 2}
	days, _ := strconv.Atoi(os.Args[1])
	
	check := 0
	for i := 0; i < days; i++ {
		check = 0
		for i := 0; i < len(toParse); i++ {
			toParse[i] = toParse[i] - 1
			if toParse[i] == -1 {
				toParse[i] = 6
				check++
			}
		}
		if check > 0 {
			for i := 0; i < check; i++ {
				toParse = append(toParse, 8)
			}
		}
		// fmt.Println(numb, "dopo il giorno:", i+1)

	}
	fmt.Println(len(toParse))
}
