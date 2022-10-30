// visita di elementi adiacenti
// quello che devo fare é controllare gli elementi minimi di una slice,
// una volta fatto posso controllare se il numero di elementi minimi
// é uguale al numero di elementi crescenti della slice

package main

import "fmt"

func main() {
	varCheck := 0
	varSlice := []int{9, 1, 3, 5, 2, 0, 8, 6}
	for i := 1; i < len(varSlice)-1; i++ {
		if varSlice[i] < varSlice[i+1] && varSlice[i] < varSlice[i-1] {
			varCheck++
		}
	}
	fmt.Println(varCheck)
}
