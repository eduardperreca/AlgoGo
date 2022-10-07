package main

import (
	"fmt"
)

func recursiveDaysLength(slice []int, check int) int {
	if check == 0 {
		return len(slice)
	}
	test := 0
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] - 1
		if slice[i] == -1 {
			slice[i] = 6
			test++
		}
	}
	if test > 0 {
		for k := 0; k < test; k++ {
			slice = append(slice, 8)
		}
	}
	fmt.Println(slice)

	return recursiveDaysLength(slice, check-1)
}

func main() {

	toParse := []int{3, 4, 3, 1, 2}
	days := 18
	t := recursiveDaysLength(toParse, days)
	fmt.Println(t)
	// 	for _, k := range numbers {
	// 		numbConv, _ := strconv.Atoi(k)
	// 		numb = append(numb, numbConv)
	// 	}

	// 	days := 256
	// 	check := 0
	// 	for i := 0; i < days ; i++ {
	// 		check = 0
	// 		for i := 0; i < len(numb); i++ {
	// 			numb[i] = numb[i] -1
	// 			if numb[i] == -1 {
	// 				numb[i] = 6
	// 				check++
	// 			}
	// 		}
	// 		if check > 0{
	// 			for i := 0; i < check; i++{
	// 				numb = append(numb, 8)
	// 			}
	// 		}
	// 		// fmt.Println(numb, "dopo il giorno:", i+1)

	// 	}
	// 	fmt.Println(len(numb))
}
