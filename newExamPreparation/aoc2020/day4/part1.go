package main

import (
	"bufio"
	"fmt"
	"os"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	cid string
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	slice := []string{}
	bigSlice := [][]string{}
	for b.Scan() {
		var dummy string
		if b.Text() != "" {
			dummy += b.Text() + " "
			slice = append(slice, dummy)
		} else {
			bigSlice = append(bigSlice, slice)
			slice = []string{}
		}
	}
	fmt.Println(bigSlice)
	for _, k := range bigSlice {
		for _, x := range k {
			fmt.Println(x)
		}
	}

}
