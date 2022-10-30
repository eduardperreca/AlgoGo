package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	numToSave := 0
	flag := false
	b := bufio.NewScanner(os.Stdin)
	for b.Scan() {
		num, _ := strconv.Atoi(b.Text())
		if num == -1 {
			break
		}

		if num > 100 {
			flag = true
			numToSave = num
		}
	}
	if flag{
		println(numToSave)
	}else{
		println("nessun numero maggiore di 100")
	}
}
