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

	for b.Scan() {
		if i != 0 {
			if len(b.Text()) > 10 {
				stringa := b.Text()
				idk := strconv.Itoa(len(stringa) - 2)
				fmt.Println(string(stringa[0]) + idk + string(stringa[len(stringa)-1]))
			} else {
				fmt.Println(b.Text())

			}
		}
		i++
	}
}
