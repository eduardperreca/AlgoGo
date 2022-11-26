package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	b := bufio.NewScanner(os.Stdin)

	i := 0
	tot := 0
	countTotal := 0
	for b.Scan() {
		tot = 0
		if i != 0 {
			for _, k := range b.Text() {
				numb, _ := strconv.Atoi(string(k))
				tot += numb
			}
			if tot > 1 {
				countTotal++
			}
		}
		i++
	}
	fmt.Println(countTotal)
}
