/* moltiplicazione somma iterata (iterative sum multiplication) ^_^	*/
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

	var risultato int
	for ; b > 0; b-- {
		risultato += a
	}
	fmt.Println("risultato: ", risultato)
	fmt.Println("sta roba ci ha messo: ", time.Since(start))
}
