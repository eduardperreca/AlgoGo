/* moltiplicazione somma iterata (iterative sum multiplication) ^_^	*/
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	var a, b, risultato int
	fmt.Println("inserisci qui due numeri da moltiplicare: ")
	fmt.Scan(&a, &b)
	for ; b > 0; b-- {
		risultato += a
	}
	fmt.Println("risultato: ", risultato)
	fmt.Println("sta roba ci ha messo: ", time.Since(start))
}
