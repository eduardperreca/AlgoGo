package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringa string
	fmt.Scan(&stringa)
	fmt.Println(strings.ToUpper(stringa[:1]) + stringa[1:])
}
