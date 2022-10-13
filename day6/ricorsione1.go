// func multiply(x, y int) int {
// 	if AAAAAA {
// 		return BBBBB
// 	} else {
// 		return x + multiply(x, y-1)
// 	}
// }

package main

import (
	"fmt"
	"os"
	"strconv"
)

// va aggiunto il caso base (y = 0) che ritorna 0

func multiply(x, y int) int{
	if y == 0 {
		return 0
	} else {
		return x + multiply(x, y-1)
	}
}

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	t := multiply(a, b)

	fmt.Println(t)

}
