package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minMax(mappa map[int]int) (int, int) {
	var min = 9999
	var max = -1
	for k, v := range mappa {
		if v < min {
			min = k
		}
		if v > max {
			max = k
		}
	}
	if mappa[min] > mappa[max] {
		return max, min
	}
	return min, max
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	random := [][]int{}
	for b.Scan() {
		line := b.Text()
		row := []int{}
		for _, c := range line {
			i, _ := strconv.Atoi(string(c))
			row = append(row, i)
		}
		random = append(random, row)
	}

	// Part 1

	var mappa = make(map[int]int)
	var listBinaryMax = []int{}
	var listBinaryMin = []int{}

	for check := 0; check < len(random[0]); check++ {
		for _, row := range random {
			for j, col := range row {
				if j == check {
					mappa[col]++
				}
			}
		}
		minimo, massimo := minMax(mappa)
		listBinaryMax = append(listBinaryMax, massimo)
		listBinaryMin = append(listBinaryMin, minimo)
		mappa = make(map[int]int)

	}

	test := strings.Trim(strings.Join(strings.Split(fmt.Sprint(listBinaryMax), " "), ""), "[]")
	test2 := strings.Trim(strings.Join(strings.Split(fmt.Sprint(listBinaryMin), " "), ""), "[]")
	dummy, _ := strconv.ParseInt(test, 2, 32)
	dummy2, _ := strconv.ParseInt(test2, 2, 32)
	fmt.Println(dummy*dummy2)

	// Part 2

	
}