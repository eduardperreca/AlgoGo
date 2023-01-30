package main

import (
	"fmt"
	"sort"
)

type Racer struct {
	start, speed int
}

type Racers []Racer

func (r Racers) Len() int {
	return len(r)
}

func (r Racers) Less(i, j int) bool {
	if r[i].start+r[i].speed*(log2N-1) == r[j].start+r[j].speed*(log2N-1) {
		return i < j
	}
	return r[i].start+r[i].speed*(log2N-1) > r[j].start+r[j].speed*(log2N-1)
}

func (r Racers) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

var log2N int

func main() {
	var N int
	fmt.Scanf("%d", &N)

	log2N = 0
	for N > 0 {
		N = N >> 1
		log2N++
	}

	racers := make(Racers, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &racers[i].start, &racers[i].speed)
	}

	sort.Sort(racers)

	fmt.Println(racers[0].start)
}
