/* moltiplicazione metodo russo (russian multiplication) ^_^	*/
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
	var risultato int
	for a > 0 {
		if a%2 == 1 {
			risultato += b
		}
		a /= 2
		b *= 2
		i++

	}
	fmt.Println("risultato: ", risultato)
	fmt.Println("ripetuto: ", i, " volte")
	fmt.Println("sta roba ci ha messo: ", time.Since(start))
}
