/*
Data da standard input una serie di interi positivi terminata da 0, stampare ’+’ ogni volta che il nuovo valore
è maggiore o uguale al precedente e ’-’ altrimenti.
*/

package main

import (
	"fmt"
)

func main() {
	var n, old int
	fmt.Scan(&n)
	for n != 0 {
		if n >= old {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
		old = n
		fmt.Scan(&n)
	}
}