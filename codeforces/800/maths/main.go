package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var stringa string
	fmt.Scan(&stringa)

	if len(stringa) == 1 {
		fmt.Println(stringa)
		return
	}

	var listInt []int
	dummy := strings.Split(stringa, "+")
	for _, k := range dummy {
		numb, _ := strconv.Atoi(k)
		listInt = append(listInt, numb)
	}
	sort.Ints(listInt)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(listInt)), "+"), "[]"))

}
