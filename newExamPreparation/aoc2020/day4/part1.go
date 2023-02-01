package main

import (
	"bufio"
	"os"
	"strings"
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
			dummy = 
			slice = append(slice, dummy)
		} else {
			bigSlice = append(bigSlice, slice)
			slice = []string{}
		}
	}

}
