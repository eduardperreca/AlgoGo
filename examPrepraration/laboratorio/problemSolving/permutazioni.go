package main

import "fmt"

type Permutazione struct {
	numero int
	nome   string
}

func main() {

	var permutazioni []Permutazione = []Permutazione{
		{1, "a"},
		{6, "f"},
		{4, "d"},
		{2, "b"},
		{3, "c"},
		{5, "e"},
		{7, "g"},
	}

	for i := 0; i < len(permutazioni); i++ {
		chiave := permutazioni[i].numero
		permutazioni[i], permutazioni[len(permutazioni)-chiave] = permutazioni[len(permutazioni)-chiave], permutazioni[i]
	}
	fmt.Println(permutazioni)
}
