
/* algoritmo potenza lineare */

/*
PSEUDO CODICE:
ALGORITMO Potenza(intero x, intero y) intero
	power <- 1
	WHILE y>0 DO
		power <- power*x
		y <- y-1
return power 

time complexity:
	3y + 3
space complexity:
	3 (x, y, power)
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	start := time.Now()

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	i := 0
	risultato := 1

	for ; b>0 ; b-- {
		risultato *= a	
		i++
	}

	fmt.Println("risultato: ", risultato)
	fmt.Println("ripetuto: ", i, " volte")
	fmt.Println("sta roba ci ha messo: ", time.Since(start))
}
