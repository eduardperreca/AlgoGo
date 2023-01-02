package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {

	rilevazioni := []int{}
	cap, intervallo := 0, 0
	fmt.Scan(&cap, &intervallo)

	b := bufio.NewScanner(os.Stdin)
	i := 0
	for b.Scan() {
		if i == cap-1 {
			break
		}
		numb, _ := strconv.Atoi(b.Text())
		rilevazioni = append(rilevazioni, numb)
		i++
	}

	tot := 0
	for i := 0; i < cap-intervallo; i++ {
		if rilevazioni[i] > rilevazioni[i+intervallo] {
			tot++
		}
	}

	fmt.Println(intervallo, tot)
}
