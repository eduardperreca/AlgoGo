package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	slices := []int{}
	matrix := [8][128]int{}
	for b.Scan() {
		array := []int{}

		for i := 0; i < 128; i++ {
			array = append(array, i)
		}
		// fmt.Println(array)
		istructions := b.Text()
		// var m int
		// binary search strange
		counti := 0
		for i := 0; i < len(istructions[:7]); i++ {
			m := len(array) / 2
			if string(istructions[i]) == "F" {
				counti = 0
				array = array[counti:m]
				// fmt.Println(array, counti, m)
			} else if string(istructions[i]) == "B" {
				counti = m
				m = len(array)
				array = array[counti:m]
				// fmt.Println(array, counti, m)
			}
		}
		row := array[0]

		//dummy implementation
		seats := []int{}
		for i := 0; i < 8; i++ {
			seats = append(seats, i)
		}
		counti = 0
		seatsIndications := istructions[7:]
		for i := 0; i < len(seatsIndications); i++ {
			m := len(seats) / 2
			if string(seatsIndications[i]) == "L" {
				counti = 0
				seats = seats[counti:m]
				// fmt.Println(seats, counti, m)
			} else if string(seatsIndications[i]) == "R" {
				counti = m
				m = len(seats)
				seats = seats[counti:m]
				// fmt.Println(array, counti, m)
			}
		}

		seat := seats[0]
		// fmt.Println(seat)
		dummyNumb := row*8 + seat
		slices = append(slices, dummyNumb)
		matrix[seat][row] = 1
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j > 5 && j < len(matrix[0])-10 && matrix[i][j] == 0 {
				fmt.Println(j*8 + i)
			}
		}
	}
	sort.Ints(slices)
	// fmt.Println(matrix)
}
