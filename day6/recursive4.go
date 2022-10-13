package main

import (
	"fmt"
	"os"
	"strconv"
)

func f_rec(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return f_rec(n-1) + f_rec(n-2)
}

func main() {
	numb, _ := strconv.Atoi(os.Args[1])
	t := f_rec(numb)
	fmt.Println(t)
}
