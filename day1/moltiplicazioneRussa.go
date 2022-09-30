/* moltiplicazione metodo russo (russian multiplication) ^_^	*/
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	var a, b int
	fmt.Println("inserisci qui due numeri da moltiplicare: ")
	fmt.Scan(&a, &b)
	var risultato int
	risultato = 0
	for a > 0 {
		if a%2 == 1 {
			risultato += b
		}
		a /= 2
		b *=  2
	}
	fmt.Println("risultato: ", risultato)
	fmt.Println("sta roba ci ha messo: ", time.Since(start))
}
