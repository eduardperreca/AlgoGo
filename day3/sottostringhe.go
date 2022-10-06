package main

import "fmt"

func main() {
	countA := 0
	countTot := 0
	for _, k := range "cdbaagbabbgbab" {
		if k == 'a'{
			countA++
			fmt.Println("somma di A: ", countA)
		}

		if k == 'b' && countA != 0 {
			countTot += countA
			fmt.Println(countTot)
		}
	}
	fmt.Println(countTot)
}
