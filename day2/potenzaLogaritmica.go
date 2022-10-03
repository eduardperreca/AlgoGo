
/* algoritmo potenza logaritmicad */

/*
PSEUDO CODICE:
AlGORITMO Potenza(intero x, intero y) intero
	if y = 0 then
		return 1
	else
		power <- Potenza(x, y/2) //  divisione intera
		power <- power * power
		if y é dispari then
			power <- power * x
		return power

// time complexity:
// T(x, y)
// se y = 0 1, 2                              costo: 2
// se y > 0 1, 2, 3, 4, 5, 7                  costo: 5
// linea 6 <= 1 volta                         costo: <= 1
// linee eseguite per calcolare (x, y/2).     costo: T(x, [y/2])
//                                            ------------------
//                                             <= 6 + T(x, [y/2])
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

	risultato, i := potenza(a, b)	

	fmt.Println("risultato: ", risultato)
	fmt.Println("ripetuto: ", i, " volte")
	fmt.Println("sta roba ci ha messo: ", time.Since(start))
}

func potenza(x, y int) (int, int) {
	if y == 0 {
		return 1, 1
	} else {
		power, i := potenza(x, y/2)
		power *= power
		if y%2 != 0 {
			power *= x
			i++
			fmt.Println("dispari: ", i)
		}
		return power, i
	}
}